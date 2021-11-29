package arisandetail

import (
	"time"

	"gorm.io/gorm"
)

type Arisan struct {
	Id       	  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt	  time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	
}
