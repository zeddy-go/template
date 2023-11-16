package main

import (
	_ "github.com/zeddy-go/database/wgorm"
	"github.com/zeddy-go/ginx"
	_ "template/config"
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
