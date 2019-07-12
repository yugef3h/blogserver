package queue

import (
	"gitlab.srgow.com/warehouse/proto/common"
	"strings"
)

type Message struct {
	id MessageID
	*common.Message
}

type MessageID string

func (id MessageID) IsTimeout() bool {
	return strings.Compare(string(id), string(TimeoutMessageID)) == 0
}

const TimeoutMessageID MessageID = "$$TIMEOUT$$"

var EmptyMessages = make([]*Message, 0, 0)
