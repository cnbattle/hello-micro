

# 基本开发流程
1. 定义接口

[hello.proto](proto/hello/hello.proto)

- 编写proto文件
```proto
syntax = "proto3";

message Request {
    string name = 1;
}

message Response {
    string msg = 1;
}

service Hello {
    rpc Hi(Request) returns (Response){
    }
}
```
- 生成`pb.go` `pb.micro.go`文件
```bash
cd proto文件所在目录
protoc --proto_path=. --micro_out=. --go_out=. [你的proto文件名].proto
```

2. 实现接口

[handler.go](srv/hello-srv/handler/handler.go)
[service.go](srv/hello-srv/service/service.go)

- 定义Handler 

3. 创建服务

[main.go](srv/hello-srv/main.go)

- NewServer服务
- Init 初始化
- 挂在接口
- 运行

# 服务参数
## 框架参数
- 1.框架API
- 2.ENV环境变量
- 3.CLI命令行参数
- 同时声明: 1<2<3
- 查找参数与变量名(--help)

`go run main.go --help`
```proto
NAME:
    - a go-micro service

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --client value                  Client for go-micro; rpc [$MICRO_CLIENT]
   --client_request_timeout value  Sets the client request timeout. e.g 500ms, 5s, 1m. Default: 5s [$MICRO_CLIENT_REQUEST_TIMEOUT]
   --client_retries value          Sets the client retries. Default: 1 (default: 1) [$MICRO_CLIENT_RETRIES]
   --client_pool_size value        Sets the client connection pool size. Default: 1 (default: 0) [$MICRO_CLIENT_POOL_SIZE]
   --client_pool_ttl value         Sets the client connection pool ttl. e.g 500ms, 5s, 1m. Default: 1m [$MICRO_CLIENT_POOL_TTL]
   --register_ttl value            Register TTL in seconds (default: 60) [$MICRO_REGISTER_TTL]
   --register_interval value       Register interval in seconds (default: 30) [$MICRO_REGISTER_INTERVAL]
   --server value                  Server for go-micro; rpc [$MICRO_SERVER]
   --server_name value             Name of the server. com.cnbattle.srv.example [$MICRO_SERVER_NAME]
   --server_version value          Version of the server. 1.1.0 [$MICRO_SERVER_VERSION]
   --server_id value               Id of the server. Auto-generated if not specified [$MICRO_SERVER_ID]
   --server_address value          Bind address for the server. 127.0.0.1:8080 [$MICRO_SERVER_ADDRESS]
   --server_advertise value        Used instead of the server_address when registering with discovery. 127.0.0.1:8080 [$MICRO_SERVER_ADVERTISE]
   --server_metadata value         A list of key-value pairs defining metadata. version=1.0.0 [$MICRO_SERVER_METADATA]
   --broker value                  Broker for pub/sub. http, nats, rabbitmq [$MICRO_BROKER]
   --broker_address value          Comma-separated list of broker addresses [$MICRO_BROKER_ADDRESS]
   --profile value                 Debug profiler for cpu and memory stats [$MICRO_DEBUG_PROFILE]
   --registry value                Registry for discovery. etcd, mdns [$MICRO_REGISTRY]
   --registry_address value        Comma-separated list of registry addresses [$MICRO_REGISTRY_ADDRESS]
   --runtime value                 Runtime for building and running services e.g local, kubernetes (default: "local") [$MICRO_RUNTIME]
   --selector value                Selector used to pick nodes for querying [$MICRO_SELECTOR]
   --transport value               Transport mechanism used; http [$MICRO_TRANSPORT]
   --transport_address value       Comma-separated list of transport addresses [$MICRO_TRANSPORT_ADDRESS]
   --help, -h                      show help
```
例子:不使用原服务名称`com.cnbattle.hello.micro.svr.hello`,用新的`com.cnbattle.hello.micro.svr.helloword`

`go run main.go --server_name=com.cnbattle.hello.micro.svr.helloword`
