package seed

import (
	"errors"
	"github.com/zeddy-go/database"
	"gorm.io/gorm"
	"template/module/test/domain"
	"template/module/test/infra/model"
)

func SeedUser(userRepository domain.IUserRepository) error {
	_, err := userRepository.First(database.Condition{"username", "test"})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userRepository.Create(&model.User{
			Username: "test",
			Password: "test",
		})
	}

	if err != nil {
		panic(err)
	}

	return nil
}
