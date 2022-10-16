package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user" gorm:"ForeignKey:UserId"`
	Photo     *Photo    `json:"photo" gorm:"ForeignKey:PhotoId"`
}

type CommentUpdate struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(comment)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
