## 介绍 

[![Build Status](https://cloud.drone.io/api/badges/cnbattle/hello-micro/status.svg)](https://cloud.drone.io/cnbattle/hello-micro)
[![LINK](https://img.shields.io/badge/link-Github-%23FF4D5B.svg?style=flat-square)](https://github.com/cnbattle/hello-micro) 
[![Go Report Card](https://goreportcard.com/badge/github.com/cnbattle/hello-micro)](https://goreportcard.com/report/github.com/cnbattle/hello-micro)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

> go-micro 相关

## TODO

### 基础服务

- [ ] 认证服务 `auth-srv`
- [ ] 用户服务 `user-srv`
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