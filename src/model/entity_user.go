package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// EntityUser 用户
type EntityUser struct {
	UserId uint64 `gorm:"primaryKey" json:"user_id"`
	UserName string `gorm:"type:varchar(255)" json:"user_name"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Salt string   `gorm:"type:varchar(255)" json:"salt"`
	NickName string `gorm:"type:varchar(255)" json:"nick_name"`
	Email string `gorm:"type:varchar(255)" json:"email"`
	Mobile string `gorm:"type:varchar(255)" json:"mobile"`
	Status int64 `gorm:"type:long" json:"status"`
	Model
}

func (EntityUser) TableName() string {
	return "user"
}

// EntityUserToken token
type EntityUserToken struct {
	UserId uint64 `gorm:"primaryKey" json:"user_id"`
	Token string `gorm:"type:varchar(255),index:unique" json:"token"`
	ExpireTime time.Time `json:"expire_time"`
	UpdateAt time.Time `json:"updated_at"`
}

func (EntityUserToken) TableName() string {
	return "user_token"
}



