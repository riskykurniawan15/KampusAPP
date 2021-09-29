package entity

import (
	"github.com/gofrs/uuid"
)

const (
	MatakuliahTableName = "matakuliah"
)

// MatakuliahModel is a model for entity.Matakuliah
type Matakuliah struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama string    `gorm:"type:varchar(200);not_null" json:"nama"`
	SKS  int       `gorm:"type:int;not_null" json:"SKS"`
	Auditable
}

func NewMatakuliah(id uuid.UUID, nama string, sks int) *Matakuliah {
	return &Matakuliah{
		ID:        id,
		Nama:      nama,
		SKS:       sks,
		Auditable: NewAuditable(),
	}
}

func (model *Matakuliah) TableName() string {
	return MatakuliahTableName
}
