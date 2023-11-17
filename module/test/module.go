package test

import (
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"template/module/test/handler"
	"template/module/test/infra/migration"
	"template/module/test/infra/repository"
)

type Module struct{}

func (m Module) RegisterRoute(r contract.IRouter) {
	container.Invoke(func(testHandler *handler.TestHandler) {
		r.GET("/test", testHandler.TestGet)
		r.POST("/test", testHandler.TestPost)
	})
}

func (m Module) Init() {
	container.Register(func() *handler.Something {
		return &handler.Something{
			B: 23,
		}
	})

	container.Register(repository.NewUserRepository)
	container.Register(handler.NewTestHandler)

	container.Invoke(migration.RegisterMigration)
}
