package main

import (
	"log"

	"github.com/cnbattle/hello-micro/srv/user-srv/handler"
	user "github.com/cnbattle/hello-micro/srv/user-srv/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("com.cnbattle.hello.micro.svr.user"),
	)

	service.Init()

	err := user.RegisterUserHandler(service.Server(), new(handler.Handler))
	if err != nil {
		log.Fatal("user.RegisterUserHandler error:", err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
