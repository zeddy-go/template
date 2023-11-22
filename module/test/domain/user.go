package domain

import (
	"github.com/zeddy-go/database"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64
	Username string
	Password string
}

type IUserRepository interface {
	database.IRepository[User, *gorm.DB]
}
