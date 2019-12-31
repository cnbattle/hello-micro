package main

import (
	"context"
	"encoding/json"
	"github.com/cnbattle/hello-micro/proto/article"
	"github.com/cnbattle/hello-micro/proto/hello"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
)

var (
	// HelloService
	helloClient hello.HelloService
	// ArticleService
	articleClient article.ArticleService
)

// Open struct
type Open struct {
}

// Fetch Fetch
func (open Open) Fetch(ctx context.Context, req *api.Request, rsq *api.Response) error {
	name := req.Get["name"].Values[0]
	articleID := req.Get["articleID"].Values[0]

	// 调用客户端
	helloRsp, err := helloClient.Hi(ctx, &hello.Request{
		Name: name,
	})
	if err != nil {
		return err
	}
	articleRsp, err := articleClient.Detail(ctx, &article.Request{
		Id: articleID,
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
		micro.Name("com.cnbattle.hello.micro.api.hello"),
	)

	// 注入
	service.Init(micro.Action(func(c *cli.Context) {
		helloClient = hello.NewHelloService("com.cnbattle.hello.micro.svr.hello", service.Client())
		articleClient = article.NewArticleService("com.cnbattle.hello.micro.svr.article", service.Client())
	}))
	// TODO

	if err := service.Run(); err != nil {
		panic(err)
	}
}
