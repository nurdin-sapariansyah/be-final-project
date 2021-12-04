package arisan

import (
	"time"

	"arisan.com/arisan/models/anggota"
	"gorm.io/gorm"
)

type Arisan struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	NamaArisan    string         `json:"namaArisan"`
	Deskripsi     string         `json:"deskripsi"`
	TotalDuit	  uint           `json:"totalDuit"`
	JumlahAnggota int            `json:"jumlahAnggota"`
	NominalIuran  int            `json:"nominalIuran"` 
	Anggota []anggota.Anggota `gorm:"many2many:arisan_anggota;" json:"anggota"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
