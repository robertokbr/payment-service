package entities

type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

func NewUser() *User {
	return &User{}
}
