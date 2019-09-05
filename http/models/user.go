package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name 		string
	Email 		string `gorm:"type:varchar(100);unique_index"`
	Phone 		string
	Password 	string
	ClientType 	string
}

type UserMapper struct {
	Id			uint		`json:"id"`
	Name 		string		`json:"name"`
	Email 		string 		`json:"email"`
	Phone 		string 		`json:"phone"`
	ClientType 	string 		`json:"client_type"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func (m User) Map() interface{} {
	memberMapper := UserMapper{
		Id:			m.ID,
		Name:       m.Name,
		Email:      m.Email,
		Phone:      m.Phone,
		ClientType: m.ClientType,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	return memberMapper
}