package main

import (
	"github.com/cnbattle/hello-micro/pkg/config"
	"log"

	proto "github.com/cnbattle/hello-micro/proto/auth"
	"github.com/cnbattle/hello-micro/srv/auth-srv/handler"
	_ "github.com/joho/godotenv/autoload"
	"github.com/micro/go-micro"
)

var (
	// microName
	microName = config.GetDefaultEnv("MICRO_NAME", "com.cnbattle.svr.auth")
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
