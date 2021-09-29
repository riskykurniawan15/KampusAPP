package entity

import (
	"github.com/gofrs/uuid"
)

const (
	NilaiTableName = "nilai"
)

// NilaiModel is a model for entity.Nilai
type Nilai struct {
	ID           uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	NilaiHuruf   string      `gorm:"type:varchar(1);not_null" json:"nilaihuruf"`
	NilaiAngka   int         `gorm:"type:int;not_null" json:"nilaiangka"`
	MahasiswaNIM *Mahasiswa  `gorm:"foreignKey:NIM" json:"mahasiswanim"`
	DosenNID     *Dosen      `gorm:"foreignKey:NID" json:"dosennid"`
	MatakuliahID *Matakuliah `gorm:"foreignKey:ID" json:"matakuliahid"`
	Auditable
}

func NewNilai(
	id uuid.UUID,
	nilaihuruf string,
	nilaiangka int,
	mahasiswanim *Mahasiswa,
	dosennid *Dosen,
	matakuliahid *Matakuliah,
) *Nilai {
	return &Nilai{
		ID:           id,
		NilaiHuruf:   nilaihuruf,
		NilaiAngka:   nilaiangka,
		MahasiswaNIM: mahasiswanim,
		MatakuliahID: matakuliahid,
		Auditable:    NewAuditable(),
	}
}

func (model *Nilai) TableName() string {
	return NilaiTableName
}
