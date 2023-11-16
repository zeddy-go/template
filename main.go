package main

import (
	"github.com/zeddy-go/ginx"
	"template/module/test"
)

type req struct {
	A int `form:"a" json:"a"`
}

func main() {

	svr := ginx.NewModule()
	svr.Register(&test.Module{})

	svr.Run()
}
