package handler

import (
	"context"
	"github.com/cnbattle/hello-micro/proto/hello"
	"github.com/cnbattle/hello-micro/srv/hello-srv/service"
)

// Handler struct
type Handler struct{}

//  Hi Hi
func (s Handler) Hi(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	rsp.Msg = service.Hi(req.Name)
	return nil
}
