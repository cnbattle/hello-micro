package main

import (
	"log"

	proto "github.com/cnbattle/hello-micro/proto/auth"
	"github.com/cnbattle/hello-micro/srv/auth-srv/handler"
	"github.com/micro/go-micro"
)

var (
	// microName
	microName = "hello.micro.svr.user.auth"
)

func main() {
	service := micro.NewService(
		micro.Name(microName),
	)

	service.Init()

	err := proto.RegisterAuthHandler(service.Server(), new(handler.Handler))
	if err != nil {
		log.Fatalf("micro %v, error:%v", microName, err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
