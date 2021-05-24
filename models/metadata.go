package models

import (
	"time"

	"gorm.io/gorm"
)

type Metadata struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"unique;type:varchar(255)"`
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Metadata) TableName() string {
	return "metadata"
}
