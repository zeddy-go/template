package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"github.com/zeddy-go/ginx"
)

type req struct {
	A int `form:"a" json:"a"`
}

type something struct {
	B int
}

type Module struct{}

func (m Module) RegisterRoute(r contract.IRouter) {
	r.GET("/test", func(ctx *gin.Context, a req) (meta ginx.IMeta, data any, err error) {
		fmt.Printf("%+v\n", a)
		return &ginx.Meta{
			Total:   20,
			PerPage: 10,
		}, gin.H{"test": true}, nil
	})

	r.POST("/test", func(a *req, b *Something) (data any, err error) {
		fmt.Printf("%+v\n", a)
		fmt.Printf("%+v\n", b)
		return gin.H{"test": true}, nil
	})
}

func (m Module) Register(sub contract.IModule) {
	panic("implement me")
}

type Something struct {
	B int
}

func (m Module) Init() {
	container.Register(func() *Something {
		return &Something{
			B: 23,
		}
	})
}
