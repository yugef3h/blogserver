package services

import (
	"context"
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
	"gitlab.srgow.com/warehouse/queue"
	"sync"
)

func Create() *SimpleMessageQueueService {
	return &SimpleMessageQueueService{
		lock:   sync.RWMutex{},
		queues: make(map[string]*queue.Queue),
	}
}

type SimpleMessageQueueService struct {
	lock   sync.RWMutex
	queues map[string]*queue.Queue
}

func (s *SimpleMessageQueueService) ReceiveMessages(ctx context.Context, in *chatroom.MessageReceiveRequest, resp *common.MessageResponse) error {
	return s.ensureQueue(in.RoomID).ReceiveMessages(ctx, in, resp)
}

func (s *SimpleMessageQueueService) SendMessages(ctx context.Context, in *chatroom.SendMessagesRequest, resp *common.Response) error {
	for _, m := range in.Messages {
		if utils.EmptyOrWhiteSpace(m.RoomID) {
			resp.Result = common.CreateErrorResult("消息中房间ID不能为空")
			return nil
		}
		id, err := utils.RandomID()
		if nil != err {
			resp.Result = common.CreateErrorResult(fmt.Sprintf("生成消息ID失败: %s", err.Error()))
			return nil
		}
		m.Id = id
	}

	for _, m := range in.Messages {
		_ = s.ensureQueue(m.RoomID).SendMessages(ctx, in, resp)
	}

	resp.Result = common.SuccessResult
	return nil
}

func (s *SimpleMessageQueueService) Shutdown() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, q := range s.queues {
		q.Shutdown()
	}
}

func (s *SimpleMessageQueueService) ensureQueue(roomID string) *queue.Queue {
	s.lock.Lock()
	defer s.lock.Unlock()

	q, ok := s.queues[roomID]
	if !ok {
		q = queue.Create(roomID, utils.EmptyString)
		s.queues[roomID] = q
	}

	return q
}
