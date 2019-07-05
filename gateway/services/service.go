package services

import (
	"context"
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/common"
	"gitlab.srgow.com/warehouse/proto/gateway"
	"strings"
	"sync"
	"sync/atomic"
)

func CreateService(
	roomService chatroom.RoomService,
	memberService chatroom.MemberService,
	receiver chatroom.MessageReceiveService,
	sender chatroom.MessageSendService) *Service {
	s := &Service{
		roomService:    roomService,
		memberService:  memberService,
		receiver:       receiver,
		sender:         sender,
		memberQueues:   make(map[string]*MemberQueue),
		roomConnectors: make(map[string]*Connector),
		lock:           sync.RWMutex{},
		exit:           &atomic.Value{},
	}
	s.exit.Store(false)
	return s
}

type Service struct {
	roomService    chatroom.RoomService
	memberService  chatroom.MemberService
	receiver       chatroom.MessageReceiveService
	sender         chatroom.MessageSendService
	memberQueues   map[string]*MemberQueue // memberID
	roomConnectors map[string]*Connector   //roomID
	lock           sync.RWMutex
	exit           *atomic.Value
}

func (s Service) SendMessage(ctx context.Context, in *gateway.SendMessageRequest, resp *common.Response) error {
	if nil == in.Request || utils.EmptyArray(in.Request.Messages) {
		resp.Result = common.CreateErrorResult("request或messages为空")
		return nil
	}
	ok, result, member := s.validateToken(ctx, in.Token)
	if !ok {
		resp.Result = result
		return nil
	}

	rResp, err := s.roomService.Available(ctx, &chatroom.AvailableRequest{MemberName: member.Name})
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	if !rResp.Result.Success {
		resp.Result = rResp.Result
		return nil
	}
	if utils.EmptyArray(rResp.Rooms) {
		resp.Result = common.CreateErrorResult("当前成员无任务房间可接收消息")
		return nil
	}

	for _, m := range in.Request.Messages {
		m.Sender = member.Name
		if !utils.Any(rResp.Rooms, func(r *chatroom.Room) bool { return strings.Compare(r.Id, m.RoomID) == 0 }) {
			resp.Result = common.CreateErrorResult(fmt.Sprint("当前用户无权在房间:", m.RoomID, "中发送消息"))
			return nil
		}
	}

	sResp, err := s.receiver.SendMessages(ctx, in.Request)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	*resp = *sResp
	return nil
}

func (s Service) ReceiveMessage(ctx context.Context, in *gateway.MessageReceiveRequest, resp *common.MessageResponse) error {
	ok, result, member := s.validateToken(ctx, in.Token)
	if !ok {
		resp.Result = result
		return nil
	}

	lResp, err := s.roomService.Available(ctx, &chatroom.AvailableRequest{MemberName: member.Name})
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}

	if !lResp.Result.Success {
		resp.Result = lResp.Result
		return nil
	}

	if utils.EmptyOrWhiteSpace(lResp.Rooms) {
		resp.Result = common.CreateErrorResult("当前成员无任何可用聊天室")
		return nil
	}

	for _, r := range lResp.Rooms {
		s.ensureConnectorAndMemberQueues(r)
	}

	resp.Messages = s.memberQueues[member.Name].Dequeue(in.ConsumeMessageIDs, in.Capacity)
	resp.Result = common.SuccessResult
	return nil
}

func (s Service) Available(ctx context.Context, in *gateway.AvailableRequest, resp *chatroom.RoomListResponse) error {
	ok, result, member := s.validateToken(ctx, in.Token)
	if !ok {
		resp.Result = result
		return nil
	}

	lResp, err := s.roomService.Available(ctx, &chatroom.AvailableRequest{MemberName: member.Name})
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	*resp = *lResp
	return nil
}

func (s Service) Login(ctx context.Context, in *chatroom.Member, resp *chatroom.LoginResponse) error {
	lResp, err := s.memberService.Login(ctx, in)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	*resp = *lResp
	return nil
}

func (s Service) Logout(ctx context.Context, in *chatroom.TokenRequest, resp *common.Response) error {
	lResp, err := s.memberService.Logout(ctx, in)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}
	*resp = *lResp
	return nil
}

func (s Service) GetMember(ctx context.Context, in *gateway.GetMemberRequest, resp *chatroom.MemberResponse) error {
	ok, result, _ := s.validateToken(ctx, in.Token)
	if !ok {
		resp.Result = result
		return nil
	}

	gResp, err := s.memberService.GetMember(ctx, in.Request)
	if nil != err {
		resp.Result = common.CreateErrorResult(err)
		return nil
	}

	*resp = *gResp
	return nil
}

func (s Service) validateToken(ctx context.Context, token string) (bool, *common.Result, *chatroom.Member) {
	vResp, err := s.memberService.Validate(ctx, &chatroom.TokenRequest{Token: token})
	if nil != err {
		return false, common.CreateErrorResult(err), nil
	}
	if !vResp.Result.Success {
		return false, vResp.Result, nil
	}

	return true, nil, vResp.Member
}

func (s *Service) Shutdown() {
	s.exit.Store(true)
}

func (s *Service) ensureConnectorAndMemberQueues(r *chatroom.Room) {
	fmt.Println("ensureConnectorAndMemberQueues:", r.Id, r.MemberNames)

	s.lock.RLock()
	defer s.lock.RUnlock()

	memberQueues := make([]*MemberQueue, 0, len(r.MemberNames))
	for _, name := range r.MemberNames {
		q, ok := s.memberQueues[name]
		if !ok {
			q = CreateMemberQueue(name)
			s.memberQueues[name] = q
		}
		memberQueues = append(memberQueues, q)
	}

	c, ok := s.roomConnectors[r.Id]
	if !ok {
		c = CreateConnector(r.Id, memberQueues, s.exit, 10, s.sender)
		s.roomConnectors[r.Id] = c
		go c.Run()
	}
}
