package test

import (
	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/contract"
	"github.com/zeddy-go/zeddy/module"
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

	return m
}

type Module struct {
	*module.BaseModule
}

func (m Module) Boot() {
	for _, item := range seed.Seeds {
		_, err := container.Invoke(item)
		if err != nil {
			panic(err)
		}
	}
}

func (m Module) RegisterRoute(r contract.IRouter, testHandler *handler.TestHandler) {
	r.GET("/test", testHandler.TestGet)
	r.POST("/test", testHandler.TestPost)
}
