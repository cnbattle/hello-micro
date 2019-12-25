package main

import (
	"github.com/cnbattle/hello-micro/proto/user"
	"github.com/cnbattle/hello-micro/srv/user-srv/handler"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.hello.micro.svr.user"),
	)

	service.Init()

	_ = user.RegisterUserHandler(service.Server(), new(handler.Handler))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
