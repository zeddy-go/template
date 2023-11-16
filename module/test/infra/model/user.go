package model

import "github.com/zeddy-go/database/wgorm"

type User struct {
	wgorm.CommonField
	Username string
	Password string
}
