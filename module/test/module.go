package test

import (
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"github.com/zeddy-go/core/module"
	"template/module/test/handler"
	"template/module/test/infra/migration"
	"template/module/test/infra/repository"
	"template/module/test/infra/seed"
)

func NewModule() *Module {
	m := &Module{
		BaseModule: module.NewBaseModule("test"),
	}

	container.Register(func() *handler.Something {
		return &handler.Something{
			B: 23,
		}
	})

	container.Register(repository.NewUserRepository)
	container.Register(handler.NewTestHandler)

	container.Invoke(migration.RegisterMigration)
	for _, item := range seed.Seeds {
		_, err := container.Invoke(item)
		if err != nil {
			panic(err)
		}
	}

	return m
}

type Module struct {
	*module.BaseModule
}

func (m Module) RegisterRoute(r contract.IRouter, testHandler *handler.TestHandler) {
	r.GET("/test", testHandler.TestGet)
	r.POST("/test", testHandler.TestPost)
}
