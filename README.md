## 介绍 

[![Build Status](https://cloud.drone.io/api/badges/cnbattle/hello-micro/status.svg)](https://cloud.drone.io/cnbattle/hello-micro)

> go-micro 相关

## TODO

### 基础服务

- [ ] 用户服务 `user-srv`
    - [ ] 登录注册服务 `login-srv`
    - [ ] 用户信息 `userinfo-srv`

- [ ] 日志服务 `log-srv`


## 微服务

[server相关](srv.md)

[client web相关](web/hello-web/main.go)

## run
```
# run micro service
~ go run srv/hello-srv/main.go
# run web service
~ cd web/hello-web/   
~ go run main.go

# test
~ curl "http://127.0.0.1:1993/hi?name=Friends"
Hello Friends
```