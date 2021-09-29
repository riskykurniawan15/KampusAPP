package entity

const (
	MahasiswaTableName = "mahasiswa"
)

// MahasiswaModel is a model for entity.Mahasiswa
type Mahasiswa struct {
	NIM     string `gorm:"type:varchar(10);primary_key" json:"nim"`
	Nama    string `gorm:"type:varchar(200);not_null" json:"name"`
	KelasID *Kelas `gorm:"foreignKey:ID" json:"kelasid"`
	Auditable
}

func NewMahasiswa(nim, nama string, kelasid *Kelas) *Mahasiswa {
	return &Mahasiswa{
		NIM:       nim,
		Nama:      nama,
		KelasID:   kelasid,
		Auditable: NewAuditable(),
	}
}

func (model *Mahasiswa) TableName() string {
	return MahasiswaTableName
}
