package models

import "time"

type SocialMedia struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name" gorm:"not null" valid:"required"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"not null" valid:"required"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           *User     `json:"user" gorm:"ForeignKey:UserId"`
}

type SocialMediaUpdate struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
