package entity

const (
	DosenTableName = "dosen"
)

// DosenModel is a model for entity.Dosen
type Dosen struct {
	NID    string `gorm:"type:varchar(10);not_null" json:"nid"`
	Nama   string `gorm:"type:varchar(200);not_null" json:"nama"`
	Kontak string `gorm:"type:varchar(15);not_null" json:"kontak"`
	Auditable
}

func NewDosen(nid, nama string, kontak string) *Dosen {
	return &Dosen{
		NID:       nid,
		Nama:      nama,
		Kontak:    kontak,
		Auditable: NewAuditable(),
	}
}

func (model *Dosen) TableName() string {
	return DosenTableName
}
