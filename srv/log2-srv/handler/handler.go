package handler

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/util/log"
)

// Handler  broker Event
func Handler(event broker.Event) error {
	log.Logf("[sub] 收到消息，请查收: %v, %v, %v", event.Message().Header, event.Message().Header["serviceName"], string(event.Message().Body))
	return nil
}
