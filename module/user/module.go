package user

import (
	"template/module/user/domain"
	"template/module/user/handler"
	"template/module/user/infra/migration"
	"template/module/user/infra/repository"
	"template/module/user/infra/seed"

	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/contract"
	"github.com/zeddy-go/zeddy/module"
)

func NewModule() *Module {
	m := &Module{
		BaseModule: module.NewBaseModule("test"),
	}

	container.Bind[*handler.Something](func() *handler.Something {
		return &handler.Something{
			B: 23,
		}
	})

	container.Bind[domain.IUserRepository](repository.NewUserRepository)
	container.Bind[*handler.TestHandler](handler.NewTestHandler)

	container.Invoke(migration.RegisterMigration)

	return m
}

type Module struct {
	*module.BaseModule
}

func (m Module) Boot() {
	for _, item := range seed.Seeds {
		err := container.Invoke(item)
		if err != nil {
			panic(err)
		}
	}
}

func (m Module) RegisterRoute(r contract.IRouter, testHandler *handler.TestHandler) {
	r.GET("/test", testHandler.TestGet)
	r.POST("/test", testHandler.TestPost)
}
