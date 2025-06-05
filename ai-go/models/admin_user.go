package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminUser struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"size:100;not null" json:"username"`
	Password  string `gorm:"size:255;not null" json:"-"`
	Role      string `gorm:"size:50;not null;default:'editor'" json:"role"`
	IsDeleted bool   `gorm:"not null;default:false" json:"is_deleted"`
	// CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`

	CreatedStr string `gorm:"-" json:"created_at"`
	UpdatedStr string `gorm:"-" json:"updated_at"`
}

// AfterFind 在查询后格式化时间
func (u *AdminUser) AfterFind(*gorm.DB) error {
	u.CreatedStr = u.CreatedAt.Format("2006-01-02 15:04:05")
	u.UpdatedStr = u.UpdatedAt.Format("2006-01-02 15:04:05")
	return nil
}

func (AdminUser) TableName() string {
	return "admin_user"
}
