package services

import (
	"context"
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
	"gitlab.srgow.com/warehouse/queue"
)

func CreateMemberQueue(memberName string) *MemberQueue {
	return &MemberQueue{
		memberName: memberName,
		q:          queue.Create(utils.EmptyString, memberName),
	}
}

type MemberQueue struct {
	memberName string
	q          *queue.Queue
}

func (mq MemberQueue) Enqueue(messages []*common.Message) {
	resp := &common.Response{}
	err := mq.q.SendMessages(context.Background(), &chatroom.SendMessagesRequest{Messages: messages}, resp)
	if nil != err {
		fmt.Println("enqueue SendMessage failed with error:", err)
	}
}

func (mq MemberQueue) Dequeue(consumeMessageIDs []string, capacity int32) []*common.Message {
	resp := &common.MessageResponse{}
	err := mq.q.ReceiveMessages(context.Background(), &chatroom.MessageReceiveRequest{
		ConsumeMessageIDs: consumeMessageIDs,
		Capacity:          capacity,
	}, resp)
	if nil != err {
		fmt.Println("dequeue ReceiveMessages failed with error:", err)
		return common.EmptyMessages
	}
	return resp.Messages
}
