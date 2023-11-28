package seed

import (
	"errors"
	"github.com/zeddy-go/zeddy/database"
	"gorm.io/gorm"
	"template/module/test/domain"
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
