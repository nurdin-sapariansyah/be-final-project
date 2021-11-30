package arisan

import (
	"time"

	"arisan.com/arisan/models/anggota"
	"gorm.io/gorm"
)

type Arisan struct {
	gorm.Model
	Id        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	NamaArisan    string         `json:"namaArisan"`
	Deskripsi     string         `json:"deskripsi"`
	TotalDuit	  uint           `json:"totalDuit"`
	JumlahAnggota int            `json:"jumlahAnggota"`
	NominalIuran  int            `json:"nominalIuran"` 
	Anggota []anggota.Anggota `gorm:"many2many:arisan_anggota;"`
}
