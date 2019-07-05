package services

import (
	"context"
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
	"sync/atomic"
	"time"
)

func CreateConnector(roomID string, memberQueues []*MemberQueue, exited *atomic.Value, capacity int, sender chatroom.MessageSendService) *Connector {
	return &Connector{
		roomID:       roomID,
		sender:       sender,
		exit:         exited,
		capacity:     capacity,
		memberQueues: memberQueues,
	}
}

type Connector struct {
	roomID       string
	sender       chatroom.MessageSendService
	exit         *atomic.Value
	capacity     int
	memberQueues []*MemberQueue
}

func (c *Connector) Run() {
	ids := make([]string, 0, 0)
	failedCount := 1
	for !c.exited() {
		resp, err := c.sender.ReceiveMessages(context.Background(), &chatroom.MessageReceiveRequest{
			RoomID:            c.roomID,
			ConsumeMessageIDs: ids,
			Capacity:          int32(c.capacity),
		})
		if nil != err {
			fmt.Println("receive room:", c.roomID, " message failed with error:", err)
			time.Sleep(time.Second * time.Duration(failedCount))
			if failedCount < 10 {
				failedCount++
			}
			continue
		}
		if !resp.Result.Success {
			fmt.Println("receive room:", c.roomID, "message failed with unsuccessfully result:", resp.Result.Error)
			time.Sleep(time.Second * time.Duration(failedCount))
			if failedCount < 10 {
				failedCount++
			}
			continue
		}
		failedCount = 1
		if !utils.EmptyArray(resp.Messages) {
			ids = utils.Select(resp.Messages, func(m *common.Message) string { return m.Id }).([]string)
			for _, q := range c.memberQueues {
				q.Enqueue(resp.Messages)
			}
		} else {
			ids = utils.EmptyStrings
		}
	}
}

func (c *Connector) exited() bool {
	return c.exit.Load().(bool)
}
