package repository

import (
	"github.com/zeddy-go/zeddy/database"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"gorm.io/gorm"
	"template/module/test/domain"
	"template/module/test/infra/model"
)

func NewUserRepository(db *wgorm.DBHolder) domain.IUserRepository {
	return &UserRepository{
		IRepository: &wgorm.Repository[model.User, domain.User]{
			IDBHolder: db,
		},
	}
}

type UserRepository struct {
	database.IRepository[domain.User, *gorm.DB]
}
