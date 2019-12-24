## 介绍 
> 实现了一个简单`go.micro.hello.micro.svr.hello`服务 `Hi`方法,简单实现了在web api中调用`go.micro.hello.micro.svr.hello `微服务

## 微服务

[server相关](srv.md)

[client web相关](web/hello-web/main.go)

## 运行
```
# 运行微服务
~ go run srv/hello-srv/main.go
# 运行web
~ cd web/hello-web/   
~ go run main.go

# test
~ curl "http://127.0.0.1:1993/hi?name=Friends"
Hello Friends
```