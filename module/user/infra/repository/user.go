package repository

import (
	"template/module/user/domain"
	"template/module/user/infra/model"

	"github.com/zeddy-go/zeddy/database"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"gorm.io/gorm"
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
