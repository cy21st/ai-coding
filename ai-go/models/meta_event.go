package models

import (
	"time"

	"gorm.io/gorm"
)

type MetaEvent struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	EventName string `gorm:"size:200;not null" json:"event_name"`
	EventDesc string `gorm:"type:text" json:"event_desc"`
	IsDeleted bool   `gorm:"not null;default:false" json:"is_deleted"`
	// CreatedAt  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
	// UpdatedAt  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
	CreatedAt  time.Time  `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt  time.Time  `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`
	Attributes []MetaAttr `gorm:"many2many:meta_relation;foreignKey:ID;joinForeignKey:EventID;References:ID;joinReferences:AttrID" json:"attributes,omitempty"`
	CreatedStr string     `gorm:"-" json:"created_at"`
	UpdatedStr string     `gorm:"-" json:"updated_at"`
}

func (m *MetaEvent) AfterFind(*gorm.DB) error {
	m.CreatedStr = m.CreatedAt.Format("2006-01-02 15:04:05")
	m.UpdatedStr = m.UpdatedAt.Format("2006-01-02 15:04:05")
	return nil
}

func (MetaEvent) TableName() string {
	return "meta_event"
}
