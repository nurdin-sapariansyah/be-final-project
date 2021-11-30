package undian

import (
	"time"

	"gorm.io/gorm"
)

type Undian struct {
	Id         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	JumlahDraw int            `json:"jumlahDraw"`
}