package global

import (
	"time"

	"gorm.io/gorm"
)

type GGG_MODEL struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
