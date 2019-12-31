package main

import (
	"github.com/cnbattle/hello-micro/proto/article"
	"github.com/cnbattle/hello-micro/srv/article-srv/handler"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("com.cnbattle.hello.micro.svr.article"),
	)

	service.Init()

	_ = article.RegisterArticleHandler(service.Server(), new(handler.Handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
