package appModel

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
