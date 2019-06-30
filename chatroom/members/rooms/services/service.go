package services

import (
	"context"
	"gitlab.srgow.com/warehouse/chatroom/rooms/manager"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
)

func Create() *Service {
	return &Service{manager.Create()}
}

type Service struct {
	roomManager *manager.Manager
}

func (s Service) Create(ctx context.Context, in *chatroom.CreateRoomRequest, resp *chatroom.CreateRoomResponse) error {
	memberNames := utils.Select(in.MemberNames, func(s string) manager.MemberName { return manager.MemberName(s) }).([]manager.MemberName)
	room, err := s.roomManager.Create(memberNames, in.Data)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	resp.Room = room.ToProto()
	return nil
}

func (s Service) Available(ctx context.Context, in *chatroom.AvailableRequest, resp *chatroom.RoomListResponse) error {
	rooms, err := s.roomManager.Available(manager.MemberName(in.MemberName))
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	resp.Rooms = utils.Select(rooms, func(r *manager.Room) *chatroom.Room { return r.ToProto() }).([]*chatroom.Room)
	return nil
}

func (s Service) Delete(ctx context.Context, in *chatroom.DeleteRoomRequest, resp *common.Response) error {
	err := s.roomManager.DeleteRoom(manager.RoomID(in.Id))
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	return nil
}
