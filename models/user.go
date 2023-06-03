package models

type User struct {
	Id       uint8  `gorm:primaryKey AutoIncrement json:"id"`
	Username string `json:"user"`
	Password string `json:"password"`
}
