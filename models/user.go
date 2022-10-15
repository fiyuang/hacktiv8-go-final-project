package models

import (
	"hacktiv8-go-final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" gorm:"not null;uniqueIndex" form:"username" valid:"required~Your username is required"`
	Email     string    `json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string    `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int32     `json:"age" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdate struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
