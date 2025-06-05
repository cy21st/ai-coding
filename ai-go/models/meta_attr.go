package models

import (
	"time"

	"gorm.io/gorm"
)

type MetaAttr struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
	// EventID   uint64    `gorm:"not null" json:"event_id"`
	AttrName  string `gorm:"size:200;not null" json:"attr_name"`
	AttrType  string `gorm:"size:50;not null;default:'string'" json:"attr_type"`
	AttrDesc  string `gorm:"type:text" json:"attr_desc"`
	IsDeleted bool   `gorm:"not null;default:false" json:"is_deleted"`
	// CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
	// UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`
	// Events    []MetaEvent `gorm:"many2many:meta_relation;" json:"events,omitempty"`
	Events     []MetaEvent `gorm:"many2many:meta_relation;foreignKey:ID;joinForeignKey:AttrID;References:ID;joinReferences:EventID" json:"events,omitempty"`
	CreatedStr string      `gorm:"-" json:"created_at"`
	UpdatedStr string      `gorm:"-" json:"updated_at"`
}

func (m *MetaAttr) AfterFind(*gorm.DB) error {
	m.CreatedStr = m.CreatedAt.Format("2006-01-02 15:04:05")
	m.UpdatedStr = m.UpdatedAt.Format("2006-01-02 15:04:05")
	return nil
}

func (MetaAttr) TableName() string {
	return "meta_attr"
}
