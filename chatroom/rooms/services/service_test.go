package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"testing"
)

func TestAvailable(t *testing.T) {
	service := Create()
	createResp := &chatroom.CreateRoomResponse{}
	err := service.Create(context.Background(), &chatroom.CreateRoomRequest{
		MemberNames: []string{"alydnh", "realgang", "tracy"},
	}, createResp)
	err = service.Create(context.Background(), &chatroom.CreateRoomRequest{
		MemberNames: []string{"alydnh", "realgang"},
	}, createResp)
	assert.Nil(t, err)

	roomsResp := &chatroom.RoomListResponse{}
	err = service.Available(context.Background(), &chatroom.AvailableRequest{
		MemberName: "alydnh",
	}, roomsResp)
	assert.Nil(t, err)
	assert.Len(t, roomsResp.Rooms, 2)

	err = service.Available(context.Background(), &chatroom.AvailableRequest{
		MemberName: "tracy",
	}, roomsResp)
	assert.Nil(t, err)
	assert.Len(t, roomsResp.Rooms, 1)
}
