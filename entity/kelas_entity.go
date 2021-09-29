package entity

import (
	"github.com/gofrs/uuid"
)

const (
	KelasTableName = "kelas"
)

// KelasModel is a model for entity.Kelas
type Kelas struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama string    `gorm:"type:varchar(200);not_null" json:"nama"`
	Auditable
}

func NewKelas(id uuid.UUID, nama string) *Kelas {
	return &Kelas{
		ID:        id,
		Nama:      nama,
		Auditable: NewAuditable(),
	}
}

func (model *Kelas) TableName() string {
	return KelasTableName
}
