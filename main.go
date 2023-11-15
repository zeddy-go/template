package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/ginx"
)

type req struct {
	A int `form:"a" json:"a"`
}

type something struct {
	B int
}

func main() {

	container.Register(func() *something {
		return &something{
			B: 23,
		}
	})

	m := gin.Default()

	m.GET("/test", ginx.GinHandler(func(ctx *gin.Context, a req) (meta ginx.IMeta, data any, err error) {
		fmt.Printf("%+v\n", a)
		return &ginx.Meta{
			Total:   20,
			PerPage: 10,
		}, gin.H{"test": true}, nil
	}))

	m.POST("/test", ginx.GinHandler(func(a req, b *something) (data any, err error) {
		fmt.Printf("%+v\n", a)
		fmt.Printf("%+v\n", b)
		return gin.H{"test": true}, nil
	}))

	m.Run()
}
