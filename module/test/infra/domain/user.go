package domain

import (
	"github.com/zeddy-go/database"
	"gorm.io/gorm"
	"template/module/test/infra/model"
)

type IUserRepository interface {
	database.IRepository[model.User, *gorm.DB]
}
