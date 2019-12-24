package main

import (
	"context"
	"encoding/json"
	"github.com/cnbattle/hello-micro/proto/article"
	"github.com/cnbattle/hello-micro/proto/hello"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
)

var (
	helloClient   hello.HelloService
	articleClient article.ArticleService
)

type Open struct {
}

func (open Open) Fetch(ctx context.Context, req *api.Request, rsq *api.Response) error {
	name := req.Get["name"].Values[0]
	articleId := req.Get["articleId"].Values[0]

	// 调用客户端
	helloRsp, err := helloClient.Hi(ctx, &hello.Request{
		Name: name,
	})
	if err != nil {
		return err
	}
	articleRsp, err := articleClient.Detail(ctx, &article.Request{
		Id: articleId,
	})
	if err != nil {
		return err
	}

	ret, err := json.Marshal(map[string]interface{}{
		"hello":   helloRsp,
		"article": articleRsp,
	})
	if err != nil {
		return err
	}
	rsq.Body = string(ret)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.hello.micro.api.hello"),
	)

	helloClient = hello.NewHelloService("go.micro.hello.micro.svr.hello", service.Client())
	articleClient = article.NewArticleService("go.micro.hello.micro.svr.article", service.Client())

	// 暴露接口
	// TODO

	if err := service.Run(); err != nil {
		panic(err)
	}
}
