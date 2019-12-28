package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cnbattle/hello-micro/proto/hello"
	logProto "github.com/cnbattle/hello-micro/proto/log"
	"github.com/cnbattle/hello-micro/proto/log2"
	"github.com/cnbattle/hello-micro/srv/hello-srv/handler"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"time"
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

// regLogger 日志Wrapper
func regLogger(cli client.Client) server.HandlerWrapper {
	// 初始化操作
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		// 中间操作
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			// log-srv
			pub := micro.NewPublisher("go.micro.hello.micro.topic.log", cli)
			_ = pub.Publish(ctx, &logProto.LogEvt{
				Msg: "Hello Friends",
			})

			// log2-srv
			event := log2.Event{
				Id:        uuid.New().String(),
				Timestamp: time.Now().Unix(),
				Message:   fmt.Sprintf("如果你看到了消息 %s, '那是因为我一直爱着你", "hello"),
			}
			body, _ := json.Marshal(event)
			_ = broker.Publish("go.micro.hello.micro.topic.event", &broker.Message{
				Header: map[string]string{
					"serviceName": "hello-srv",
				},
				Body: body,
			})
			return handlerFunc(ctx, req, rsp)
		}
	}
}
