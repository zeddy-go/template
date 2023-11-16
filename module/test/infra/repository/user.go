package repository

import (
	"github.com/zeddy-go/database"
	"github.com/zeddy-go/database/wgorm"
	"gorm.io/gorm"
	"template/module/test/infra/domain"
	"template/module/test/infra/model"
)

func NewUserRepository(db *wgorm.DBHolder) domain.IUserRepository {
	return &UserRepository{
		IRepository: &wgorm.Repository[model.User]{
			IDBHolder: db,
		},
	}
}

type UserRepository struct {
	database.IRepository[model.User, *gorm.DB]
}
