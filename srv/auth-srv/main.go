package main

import (
	"github.com/cnbattle/hello-micro/pkg/config"
	"log"

	proto "github.com/cnbattle/hello-micro/proto/auth"
	"github.com/cnbattle/hello-micro/srv/auth-srv/handler"
	_ "github.com/joho/godotenv/autoload"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var (
	// microName
	microName = config.GetDefaultEnv("MICRO_NAME", "com.cnbattle.svr.auth")
)

func main() {
	// etcdv3 registerDrive
	registerDrive := etcdv3.NewRegistry(func(options *registry.Options) {
		// etcd 地址
		options.Addrs = []string{"127.0.0.1:2379"}
		// etcd 用户名密码,如果设置的话
		etcdv3.Auth("root", "password")(options)
	})
	service := micro.NewService(
		micro.Name(microName),
		micro.Registry(registerDrive),
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
