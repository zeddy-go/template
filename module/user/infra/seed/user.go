package seed

import (
	"errors"
	"template/module/user/domain"

	"github.com/zeddy-go/zeddy/database"
	"gorm.io/gorm"
)

func SeedUser(userRepository domain.IUserRepository) error {
	_, err := userRepository.First(database.Condition{"username", "test"})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userRepository.Create(&domain.User{
			Username: "test",
			Password: "test",
		})
	}

	if err != nil {
		panic(err)
	}

	return nil
}
