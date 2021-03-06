package main

import (
	"context"

	logProto "github.com/cnbattle/hello-micro/proto/log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

// Sub struct
type Sub struct {
}

// Process 处理日志
func (s Sub) Process(ctx context.Context, evt *logProto.LogEvt) error {
	// 业务逻辑
	log.Logf("[sub] 收到日志: %v", evt.Msg)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.cnbattle.hello.micro.srv.log"),
		micro.Version("latest"),
	)

	service.Init()

	_ = micro.RegisterSubscriber("com.cnbattle.hello.micro.topic.log", service.Server(), &Sub{})

	if err := service.Run(); err != nil {
		panic(err)
	}
}
