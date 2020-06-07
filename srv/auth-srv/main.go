package main

import (
	"log"

	"github.com/cnbattle/hello-micro/pkg/config"
	proto "github.com/cnbattle/hello-micro/proto/auth"
	"github.com/cnbattle/hello-micro/srv/auth-srv/handler"
	_ "github.com/joho/godotenv/autoload"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/transport"

	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/transport/nats"
)

var (
	// microName
	microName = config.GetDefaultEnv("MICRO_NAME", "com.cnbattle.svr.auth")
)

func main() {
	// etcdv3 registerDrive
	registerDrive := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})

	// rabbitmq brokerDrive
	brokerDrive := rabbitmq.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"127.0.0.1:5672"}
	})

	// nats transportDrive
	transportDrive := nats.NewTransport(func(options *transport.Options) {
		options.Addrs = []string{"127.0.0.1:4222"}
	})

	service := micro.NewService(
		micro.Name(microName),
		micro.Registry(registerDrive),
		micro.Broker(brokerDrive),
		micro.Transport(transportDrive),
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
