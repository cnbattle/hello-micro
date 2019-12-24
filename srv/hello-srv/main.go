package main

import (
	"context"
	"encoding/json"

	"github.com/cnbattle/hello-micro/proto/hello"
	logProto "github.com/cnbattle/hello-micro/proto/log"
	"github.com/cnbattle/hello-micro/srv/hello-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
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
			evt := logProto.LogEvt{
				Msg: "hello",
			}
			body, err := json.Marshal(evt)
			if err != nil {
				log.Info("[hello] json.Marshal err:", err)
			}

			msg := &broker.Message{
				Header: map[string]string{
					"serviceName": "hello",
				},
				Body: body,
			}

			if err := pub.Publish(ctx, msg); err != nil {
				log.Info("[hello] 发送日志 err:", err)
			}

			return handlerFunc(ctx, req, rsp)
		}
	}
}
