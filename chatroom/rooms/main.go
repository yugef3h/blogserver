package main

import (
	"github.com/micro/go-micro"
	"gitlab.srgow.com/warehouse/chatroom/rooms/services"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"time"
)

func main() {
	serviceName := utils.GetMicroServiceName("room")
	service := micro.NewService(
		micro.Name(serviceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	roomService := services.Create()
	err := chatroom.RegisterRoomServiceHandler(service.Server(), roomService)
	if nil != err {
		panic(err)
	}

	err = service.Run()
	if nil != err {
		panic(err)
	}
}
