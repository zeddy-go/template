package main

import (
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/module"
	"github.com/zeddy-go/database/migrate"
	"github.com/zeddy-go/database/wgorm"
	"github.com/zeddy-go/ginx"
	"template/config"
	"template/module/test"
)

type req struct {
	A int `form:"a" json:"a"`
}

func main() {
	module.Init(
		&config.Module{},
		&wgorm.Module{},
		&migrate.Module{},
	)

	svr := ginx.NewModule()
	svr.Register(&test.Module{})

	container.Invoke(func(m migrate.IMigrator) {
		err := m.Migrate()
		if err != nil {
			panic(err)
		}
	})

	svr.Run()
}
