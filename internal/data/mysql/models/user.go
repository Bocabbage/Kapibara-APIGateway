package models

import "time"

type User struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Account    string    `gorm:"column:account;unique;not null"`
	UserName   string    `gorm:"column:user_name;unique;not null"`
	PwdHash    string    `gorm:"column:pwd_hash;not null"`
	CreateAt   time.Time `gorm:"column:create_time;autoCreateTime;not null"`
	UpdateAt   time.Time `gorm:"column:update_time;autoUpdateTime;not null"`
	Status     int       `gorm:"column:status;default:0"`
	RoleBitmap int64     `gorm:"column:role_bitmap;default:1"`
}
