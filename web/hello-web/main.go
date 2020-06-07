package main

import (
	"context"
	"github.com/cnbattle/hello-micro/proto/hello"
	"github.com/micro/go-micro/v2/web"
	"net/http"
)

var (
	// srvClient 声明服务
	srvClient hello.HelloService
)

func main() {
	service := web.NewService(
		web.Name("com.cnbattle.hello.micro.web.hello"),
		web.Address(":1993"),
		web.StaticDir("html"),
	)
	_ = service.Init()
	// 实例化服务
	srvClient = hello.NewHelloService("com.cnbattle.hello.micro.svr.hello", service.Options().Service.Client())
	service.HandleFunc("/hi", Hi)

	if err := service.Run(); err != nil {
		panic(err)
	}
}

// Hi Hi
func Hi(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	req := &hello.Request{
		Name: name,
	}
	// 调用服务
	rsp, err := srvClient.Hi(context.Background(), req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(rsp.Msg))
}
