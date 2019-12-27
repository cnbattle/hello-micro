package main

import (
	"log"

	proto "github.com/cnbattle/hello-micro/proto/login"
	"github.com/micro/go-micro"
)

var (
	microName = "hello.micro.svr.user.login"
)

func main() {
	service := micro.NewService(
		micro.Name(microName),
	)

	service.Init()

	err := proto.RegisterLoginHandler(service.Server(), new(handler))
	if err != nil {
		log.Fatalf("micro %v, error:%v", microName, err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
