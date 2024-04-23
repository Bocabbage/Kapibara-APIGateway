package mysql

import "time"

type Users struct {
	AccountID  int       `gorm:"column:account_id;primary_key;AUTO_INCREMENT"`
	Account    string    `gorm:"column:account;unique;NOT NULL"`
	UserName   string    `gorm:"column:user_name;NOT NULL"`
	PwdHash    string    `gorm:"column:pwd_hash;NOT NULL"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	Status     int       `gorm:"column:status;default:0"` // 0: not-in-use, 1: not-active, 2: in-use, 3: black-list
	RoleBitmap int64     `gorm:"column:role_bitmap;default:1"`
}
