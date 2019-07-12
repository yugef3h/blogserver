package queue

import (
	"context"
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/common/workqueue"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
	"sync"
	"time"
)

func Create(roomID, memberID string) *Queue {
	return &Queue{
		roomID:     roomID,
		memberID:   memberID,
		lock:       sync.RWMutex{},
		queue:      workqueue.New(),
		messages:   make(map[MessageID]*Message),
		processing: make(map[MessageID]bool),
	}
}

type Queue struct {
	roomID     string
	memberID   string
	lock       sync.RWMutex
	queue      workqueue.Interface
	messages   map[MessageID]*Message
	processing map[MessageID]bool
}

func (s Queue) RoomID() string {
	return s.roomID
}

func (s *Queue) Shutdown() {
	s.queue.ShutDown()
}

func (s *Queue) ReceiveMessages(ctx context.Context, in *chatroom.MessageReceiveRequest, resp *common.MessageResponse) error {
	messages := s.consume(in.ConsumeMessageIDs)
	if !utils.EmptyArray(messages) {
		resp.Result = common.SuccessResult
		resp.Messages = utils.Select(messages, func(m *Message) *common.Message { return m.Message }).([]*common.Message)
		return nil
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*60))
	mChan := make(chan []MessageID)
	defer close(mChan)
	go func() {
		l := utils.MaxInt(1, int(in.Capacity))
		l = utils.MinInt(s.queue.Len(), l)
		if l <= 0 {
			l = 1
		}
		ids := make([]MessageID, 0, l)
		for len(ids) < l {
			item, shutdown := s.queue.Get()
			fmt.Println("dequeue:", item, shutdown)
			if shutdown {
				break
			}
			if id, ok := item.(MessageID); ok {
				if id.IsTimeout() {
					s.queue.Done(item)
					return
				}
				ids = append(ids, id)
			}
			s.queue.Done(item)
		}
		mChan <- ids
	}()

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("receive done.")
				s.queue.Add(TimeoutMessageID)
				resp.Messages = common.EmptyMessages
				resp.Result = common.SuccessResult
				return nil
			}
		case ids := <-mChan:
			if !utils.EmptyArray(ids) {
				fmt.Println("message received:", ids)
				cancel()
				resp.Messages = s.publish(ids)
				resp.Result = common.SuccessResult
				return nil
			}
		}
	}
}

func (s *Queue) SendMessages(ctx context.Context, in *chatroom.SendMessagesRequest, resp *common.Response) error {
	messages := make([]*Message, 0, len(in.Messages))
	for _, m := range in.Messages {
		if utils.EmptyOrWhiteSpace(m.Id) {
			id, err := utils.RandomID()
			m.Id = id
			if nil != err {
				resp.Result = common.CreateErrorResult(fmt.Sprintf("生成消息ID失败: %s", err.Error()))
				return nil
			}
		}
		if utils.EmptyOrWhiteSpace(m.Timestamp) {
			m.Timestamp = utils.ToDatetimeStringWithoutDash(time.Now())
		}
		messages = append(messages, &Message{id: MessageID(m.Id), Message: m})
	}
	for _, m := range messages {
		s.queue.Add(m.id)
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	for _, m := range messages {
		s.messages[m.id] = m
	}

	resp.Result = common.SuccessResult
	return nil
}

func (s *Queue) publish(ids []MessageID) []*common.Message {
	s.lock.RLock()
	defer s.lock.RUnlock()
	messages := make([]*common.Message, 0, len(ids))
	for _, id := range ids {
		m, ok := s.messages[id]
		if ok {
			messages = append(messages, m.Message)
			s.processing[id] = true
		}
	}
	return messages
}

func (s *Queue) consume(ids []string) []*Message {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, id := range ids {
		delete(s.processing, MessageID(id))
		delete(s.messages, MessageID(id))
	}

	if len(s.processing) > 0 {
		messages := make([]*Message, 0, len(s.processing))
		for id := range s.processing {
			if m, ok := s.messages[id]; ok {
				messages = append(messages, m)
			}
		}

		return messages
	}

	return EmptyMessages
}
