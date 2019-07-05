package main

import (
	"github.com/micro/go-micro"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/gateway/services"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"gitlab.srgow.com/warehouse/proto/gateway"
	"time"
)

func main() {
	serviceName := utils.GetMicroServiceName("api")
	service := micro.NewService(
		micro.Name(serviceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	serviceName = utils.GetMicroServiceName("room")
	roomService := chatroom.NewRoomService(serviceName, service.Client())

	serviceName = utils.GetMicroServiceName("member")
	memberService := chatroom.NewMemberService(serviceName, service.Client())

	serviceName = utils.GetMicroServiceName("queue")
	sender := chatroom.NewMessageSendService(serviceName, service.Client())
	receiver := chatroom.NewMessageReceiveService(serviceName, service.Client())

	gatewayService := services.CreateService(roomService, memberService, receiver, sender)

	err := gateway.RegisterRoomHandler(service.Server(), gatewayService)
	if nil != err {
		panic(err)
	}

	err = gateway.RegisterChatHandler(service.Server(), gatewayService)
	if nil != err {
		panic(err)
	}

	err = gateway.RegisterMemberHandler(service.Server(), gatewayService)
	if nil != err {
		panic(err)
	}

	defer gatewayService.Shutdown()
	if err := service.Run(); nil != err {
		panic(err)
	}
}
