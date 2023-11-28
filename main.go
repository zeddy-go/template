package main

import (
	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"github.com/zeddy-go/zeddy/http/ginx"
	_ "template/config"
	"template/module/test"
)

func main() {
	svr := ginx.NewModule()

	svr.Register(
		wgorm.NewModule(),
		migrate.NewModule(),
		test.NewModule(),
	)

	svr.Boot()

	svr.Serve()
}
