package models

import (
	"time"

	"gorm.io/gorm"
)

type MetaRelation struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	EventID   uint64 `gorm:"not null" json:"event_id"`
	AttrID    uint64 `gorm:"not null" json:"attr_id"`
	IsDeleted bool   `gorm:"not null;default:false" json:"is_deleted"`
	// CreatedAt  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"-"`
	CreatedStr string    `gorm:"-" json:"created_at"`
	Event      MetaEvent `gorm:"foreignKey:EventID" json:"event"`
	Attr       MetaAttr  `gorm:"foreignKey:AttrID" json:"attr"`
}

func (m *MetaRelation) AfterFind(*gorm.DB) error {
	m.CreatedStr = m.CreatedAt.Format("2006-01-02 15:04:05")
	return nil
}

func (MetaRelation) TableName() string {
	return "meta_relation"
}
