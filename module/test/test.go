package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"github.com/zeddy-go/ginx"
	"template/module/test/infra/domain"
	"template/module/test/infra/migration"
	"template/module/test/infra/repository"
)

type req struct {
	A int `form:"a" json:"a"`
}

type something struct {
	B int
}

type Module struct{}

func (m Module) RegisterRoute(r contract.IRouter) {
	r.GET("/test", func(ctx *gin.Context, a req, repository domain.IUserRepository) (meta ginx.IMeta, data any, err error) {
		//fmt.Printf("%+v\n", a)
		list, err := repository.List()
		if err != nil {
			panic(err)
		}
		println(len(list))
		fmt.Printf("%+v\n", list[0])
		return &ginx.Meta{
			Total:   20,
			PerPage: 10,
		}, gin.H{"test": true}, nil
	})

	r.POST("/test", func(a *req, b *Something) (data any, err error) {
		fmt.Printf("%+v\n", a)
		fmt.Printf("%+v\n", b)
		return gin.H{"test": true}, nil
	})
}

func (m Module) Register(subs ...contract.IModule) {
	panic("implement me")
}

type Something struct {
	B int
}

func (m Module) Init() {
	container.Register(func() *Something {
		return &Something{
			B: 23,
		}
	})

	container.Register(repository.NewUserRepository)

	container.Invoke(migration.RegisterMigration)
}
