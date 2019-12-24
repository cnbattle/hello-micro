package main

import (
	"context"

	logProto "github.com/cnbattle/hello-micro/proto/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

type Sub struct {
}

func (s Sub) Process(ctx context.Context, evt *logProto.LogEvt) error {
	// 业务逻辑
	log.Logf("[sub] 收到日志: %v", evt.Msg)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.hello.micro.srv.log"),
		micro.Version("latest"),
	)

	service.Init()

	_ = micro.RegisterSubscriber("go.micro.hello.micro.topic.log", service.Server(), &Sub{})

	if err := service.Run(); err != nil {
		panic(err)
	}
}
