package hello

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linghechen/go-starter/pkg/framework/web"
	"github.com/linghechen/go-starter/pkg/xerr"
)

type HelloController struct {
}

func NewHelloHandler() web.Controller {
	return &HelloController{}
}

func (h *HelloController) Registry(r *gin.Engine) {
	group := r.Group("/api/v1/hello")
	{
		group.GET("/", web.Wrap(h.GetHello))
	}
}

func (h *HelloController) GetHello(c context.Context, req *HelloReq) (*HelloResp, xerr.Error) {
	return &HelloResp{
		Reply: "hello " + req.Name,
		Time:  time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
