package test

import (
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"github.com/zeddy-go/core/module"
	"template/module/test/handler"
	"template/module/test/infra/migration"
	"template/module/test/infra/repository"
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

	return m
}

type Module struct {
	*module.BaseModule
}

func (m Module) RegisterRoute(r contract.IRouter, testHandler *handler.TestHandler) {
	r.GET("/test", testHandler.TestGet)
	r.POST("/test", testHandler.TestPost)
}
