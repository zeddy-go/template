package main

import (
	"github.com/zeddy-go/core/container"
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

	container.Invoke(func(m migrate.IMigrator) {
		err := m.Migrate()
		if err != nil {
			panic(err)
		}
	})

	svr.Run()
}
