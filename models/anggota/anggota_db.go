package anggota

import (
	"time"

	"gorm.io/gorm"
)

type Anggota struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Nama      string         `json:"name"`
	NomorHp   string         `json:"nomorHp"`
	Password  string         `json:"password"`
	Pekerjaan string		 `json:"pekerjaan"`
}
