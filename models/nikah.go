package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Create_Nikah/config"
	)

var db *gorm.DB

type Nikah struct {
	gorm.Model
	NamaMempelaiLaki string `json:"nama_mempelai_laki"`
	AlamatMempelaiLaki string `json:"alamat_mempelai_laki"`
	TanggalPernikahan string `json:"tanggal_pernikahan"`
	NamaAyahMempelaiLaki string `json:"nama_ayah_mempelai_laki"`
	NamaIbuMempelaiLaki string `json:"nama_ibu_mempelai_laki"`
	NamaLengkapMempelaiPerempuan string `json:"nama_lengkap_mempelai_perempuan"`
	AlamatMempelaiPerempuan string `json:"alamat_mempelai_perempuan"`
	NamaAyahMempelaiPerempuan string `json:"ayah_mempelai_perempuan"`
	NamaIbuMempelaiPerempuan string `json:"nama_ibu_mempelai_perempuan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Nikah{})
}

func (b *Nikah) CreateNikah() *Nikah {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllNikahs() []Nikah {
	var Nikahs []Nikah
	db.Find(&Nikahs)
	return Nikahs
}

func GetNikahbyId(id int64) (*Nikah, *gorm.DB) {
	var getNikah Nikah
	db := db.Where("ID=?", id).Find(&getNikah)
	return &getNikah, db
}

func DeleteNikah(id int64) Nikah {
	var nikah Nikah
	db.Where("ID=?", id).Delete(nikah)
	return nikah 
}