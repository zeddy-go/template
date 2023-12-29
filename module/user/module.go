package user

import (
	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/contract"
	"github.com/zeddy-go/zeddy/module"
	"template/module/user/domain"
	"template/module/user/handler"
	"template/module/user/infra/migration"
	"template/module/user/infra/repository"
	"template/module/user/infra/seed"
)

func NewModule() *Module {
	m := &Module{
		BaseModule: module.NewBaseModule("test"),
	}

	return m
}

type Module struct {
	*module.BaseModule
}

func (m Module) Init() (err error) {
	err = container.Bind[*handler.Something](func() *handler.Something {
		return &handler.Something{
			B: 23,
		}
	})
	if err != nil {
		return
	}

	err = container.Bind[domain.IUserRepository](repository.NewUserRepository)
	if err != nil {
		return
	}

	err = container.Bind[*handler.TestHandler](handler.NewTestHandler)
	if err != nil {
		return
	}

	err = container.Invoke(migration.RegisterMigration)
	if err != nil {
		return
	}

	return
}

func (m Module) Boot() (err error) {
	for _, item := range seed.Seeds {
		err = container.Invoke(item)
		if err != nil {
			return
		}
	}

	err = container.Invoke(func(r contract.IRouter, testHandler *handler.TestHandler) {
		r.GET("/test", testHandler.TestGet)
		r.POST("/test", testHandler.TestPost)
	})
	if err != nil {
		return
	}

	return
}
