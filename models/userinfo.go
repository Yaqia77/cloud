package models

import (
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	UserID      uint      `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username"  gorm:"not null"`
	Password    string    `json:"password"  gorm:"not null"`
	Email       string    `json:"email" gorm:"uniqueIndex;size:255" binding:"required"`
	Avatar      string    `json:"avatar" gorm:"type:text"`
	Token       string    `json:"token" `
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"autoUpdateTime"`
	Status      int       `json:"status" gorm:"type:tinyint(1)"`
	UseSpace    int       `json:"use_space" gorm:"type:int(11)"`
	MaxSpace    int       `json:"max_space" gorm:"type:int(11)"`
}

type UserInfoResponse struct {
	UserID      uint      `json:"user_id" `
	Username    string    `json:"username" `
	Email       string    `json:"email" `
	Avatar      string    `json:"avatar" `
	Token       string    `json:"token" `
	CreatedAt   time.Time `json:"created_at" `
	LastLoginAt time.Time `json:"last_login_at" `
	Status      int       `json:"status"`
	UseSpace    int       `json:"use_space"`
	MaxSpace    int       `json:"max_space"`
}

func AutoMigrateUserInfoTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserInfo{}, &UserInfoResponse{}); err != nil {
		return err
	}
	return nil
}
