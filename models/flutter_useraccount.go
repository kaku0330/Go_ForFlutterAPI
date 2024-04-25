package models

type User struct {
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}
