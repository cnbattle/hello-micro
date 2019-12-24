package main

import (
	"context"
	"github.com/cnbattle/hello-micro/proto/hello"
	logProto "github.com/cnbattle/hello-micro/proto/log"
	"github.com/cnbattle/hello-micro/srv/hello-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.hello.micro.svr.hello"),
	)

	service.Init(
		micro.WrapHandler(
			regLogger(service.Client()),
		),
		micro.BeforeStart(func() error {
			log.Log("启动前的日志")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log("启动后的日志")
			return nil
		}),
		//... 其他设置
	)

	_ = hello.RegisterHelloHandler(service.Server(), new(handler.Handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// 日志Wrapper
func regLogger(cli client.Client) server.HandlerWrapper {
	//初始化操作
	pub := micro.NewPublisher("go.micro.hello.micro.topic.log", cli)
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		// 中间操作
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			_ = pub.Publish(ctx, &logProto.LogEvt{
				Msg: "hello",
			})
			return handlerFunc(ctx, req, rsp)
		}
	}
}
