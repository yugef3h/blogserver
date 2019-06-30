package main

import (
	"github.com/micro/go-micro"
	"gitlab.srgow.com/warehouse/chatroom/receiver/services"
	"gitlab.srgow.com/warehouse/common/utils"
	"gitlab.srgow.com/warehouse/proto/chatroom"
	"time"
)

func main() {
	serviceName := utils.GetMicroServiceName("queue")
	service := micro.NewService(
		micro.Name(serviceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	receiverService := services.Create()
	err := chatroom.RegisterMessageReceiveServiceHandler(service.Server(), receiverService)
	if nil != err {
		panic(err)
	}

	err = chatroom.RegisterMessageSendServiceHandler(service.Server(), receiverService)
	if nil != err {
		panic(err)
	}

	err = service.Run()
	if nil != err {
		panic(err)
	}
}
