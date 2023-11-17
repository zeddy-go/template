package main

import (
	"github.com/zeddy-go/database/migrate"
	"github.com/zeddy-go/database/wgorm"
	"github.com/zeddy-go/ginx"
	"template/config"
	"template/module/test"
)

func main() {
	svr := ginx.NewModule()
	svr.Register(
		&config.Module{},
		&wgorm.Module{},
		&migrate.Module{},
		&test.Module{},
	)

	svr.Boot()

	svr.Serve()
}
