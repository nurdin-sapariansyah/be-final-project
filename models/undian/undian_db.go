package undian

import (
	"time"

	"arisan.com/arisan/models/arisan"
	"gorm.io/gorm"
)

type Undian struct {
	Id         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	JumlahDraw int            `json:"jumlahDraw"`
	ArisanRefer int			  `json:"arisan"`
	Arisan      arisan.Arisan `gorm:"foreignKey:ArisanRefer"`
}