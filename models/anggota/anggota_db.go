package anggota

import (
	"time"

	"gorm.io/gorm"
)

type Anggota struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `json:"nama"`
	Email   string         `json:"email"`
	Password  string         `json:"password"`
	Pekerjaan string		 `json:"pekerjaan"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
