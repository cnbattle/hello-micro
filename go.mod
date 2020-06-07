module github.com/cnbattle/hello-micro

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/silenceper/wechat v1.2.6
)
