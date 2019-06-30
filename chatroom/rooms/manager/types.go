package manager

import (
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
)

type Room struct {
	id      RoomID
	members map[MemberName]bool
	data    map[string]string
}

func (r Room) ToProto() *chatroom.Room {
	memberNames := utils.MapKeys(r.members).([]MemberName)
	return &chatroom.Room{
		Id:          string(r.id),
		MemberNames: utils.Select(memberNames, func(m MemberName) string { return string(m) }).([]string),
		Data:        r.data,
	}
}

type RoomID string

type MemberName string

var EmptyRooms = make([]*Room, 0, 0)
