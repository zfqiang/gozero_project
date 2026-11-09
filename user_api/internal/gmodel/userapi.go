package gmodel

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}
