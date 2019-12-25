package handler

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func Handler(event broker.Event) error {
	log.Logf("[sub] 收到消息，请查收: %v, %v, %v", event.Message().Header, event.Message().Header["serviceName"], string(event.Message().Body))
	return nil
}
