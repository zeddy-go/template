package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zeddy-go/ginx"
)

func main() {
	r := gin.Default()

	r.GET("test", ginx.GinHandler(func(ctx *gin.Context) {}))

	r.Run()
}
