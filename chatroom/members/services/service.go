package services

import (
	"context"
	"gitlab.srgow.com/warehouse/chatroom/members/manager"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
)

func Create() *Service {
	return &Service{memberManager: manager.Create()}
}

type Service struct {
	memberManager *manager.Manager
}

func (s Service) GetMember(ctx context.Context, in *chatroom.GetMemberRequest, resp *chatroom.MemberResponse) error {
	m, err := s.memberManager.GetMember(manager.MemberName(in.MemberName))
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	resp.Member = m.ToProto()
	return nil
}

func (s Service) Validate(ctx context.Context, in *chatroom.TokenRequest, resp *chatroom.MemberResponse) error {
	m, err := s.memberManager.Validate(manager.Token(in.Token))
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	resp.Member = m.ToProto()
	return nil
}


// login token refresh
func (s Service) Login(ctx context.Context, in *chatroom.Member, resp *chatroom.LoginResponse) error {
	m, err := s.memberManager.Login(manager.MemberName(in.Name), in.Data)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}

	resp.Token = string(m.Token())
	resp.Result = common.SuccessResult
	return nil
}

func (s Service) Logout(ctx context.Context, in *chatroom.TokenRequest, resp *common.Response) error {
	err := s.memberManager.Logout(manager.Token(in.Token))
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	resp.Result = common.SuccessResult
	return nil
}

func (s Service) SetData(context.Context, *chatroom.SetDataRequest, *common.Response) error {
	panic("implement me")
}
