package pembayaran

import (
	"time"

	"gorm.io/gorm"
)

type Pembayaran struct {
	Id       	  uint           `gorm:"primaryKey" json:"id"`
	StatusPembayaran bool		 `json:"statusPembayaran"`
	StatusPenjatahan bool		 `json:"statusPenjatahan"`
	CreatedAt	  time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
