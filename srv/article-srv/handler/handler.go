package handler

import (
	"context"

	"github.com/cnbattle/hello-micro/proto/article"
	"github.com/cnbattle/hello-micro/srv/article-srv/service"
)

type Handler struct{}

func (s Handler) Detail(ctx context.Context, req *article.Request, rsp *article.Response) error {
	detail := service.Detail(req.Id)
	rsp.Id = detail.Id
	rsp.UID = detail.UID
	rsp.Title = detail.Title
	rsp.Content = detail.Content
	rsp.ViewNum = detail.ViewNum
	return nil
}
