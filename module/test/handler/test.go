package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeddy-go/zeddy/http/ginx"
	"template/module/test/domain"
)

func NewTestHandler(userRepository domain.IUserRepository) *TestHandler {
	return &TestHandler{
		userRepository: userRepository,
	}
}

type TestHandler struct {
	userRepository domain.IUserRepository
}

func (t *TestHandler) TestGet(ctx *gin.Context, a req) (meta ginx.IMeta, data any, err error) {
	//fmt.Printf("%+v\n", a)
	list, err := t.userRepository.List()
	if err != nil {
		panic(err)
	}
	println(len(list))
	if len(list) > 0 {
		fmt.Printf("%+v\n", list[0])
	}
	return &ginx.Meta{
		Total:   20,
		PerPage: 10,
	}, gin.H{"test": true}, nil
}

func (t *TestHandler) TestPost(a *req, b *Something) (data any, err error) {
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	return gin.H{"test": true}, nil
}
