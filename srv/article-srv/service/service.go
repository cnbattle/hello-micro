package service

import "github.com/cnbattle/hello-micro/proto/article"

func Detail(articleId string) (rsp article.Response) {
	// ... real code
	rsp.Id = articleId
	rsp.Uid = "123"
	rsp.Title = "demo"
	rsp.Content = "content"
	rsp.ViewNum = 6302
	return
}
