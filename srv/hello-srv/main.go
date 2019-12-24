package main

import (
	"github.com/cnbattle/hello-micro/proto/hello"
	"github.com/cnbattle/hello-micro/srv/hello-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.hello.micro.svr.hello"),
	)

	service.Init(
		micro.BeforeStart(func() error {
			log.Log("启动前")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log("启动后")
			return nil
		}),
		//... 其他设置
	)

	_ = hello.RegisterHelloHandler(service.Server(), new(handler.Handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
