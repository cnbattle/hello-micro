package service

import "github.com/cnbattle/hello-micro/proto/article"

// Detail article Detail
func Detail(articleID string) (rsp article.Response) {
	// ... real code
	rsp.Id = articleID
	rsp.UID = "123"
	rsp.Title = "demo"
	rsp.Content = "content"
	rsp.ViewNum = 6302
	return
}
