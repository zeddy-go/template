package main

import (
	_ "template/config"
	"template/module/user"

	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"github.com/zeddy-go/zeddy/http/ginx"
)

func main() {
	svr := ginx.NewModule()

	svr.Register(
		wgorm.NewModule(),
		migrate.NewModule(),
		user.NewModule(),
	)

	svr.Boot()

	svr.Serve()
}
