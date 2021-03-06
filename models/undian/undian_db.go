package undian

import (
	"time"

	"arisan.com/arisan/models/arisan"
	"gorm.io/gorm"
)

type Undian struct {
	Id         uint           `gorm:"primaryKey" json:"id"`
	JumlahDraw int            `json:"jumlahDraw"`
	Arisan     arisan.Arisan `json:"arisan"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}