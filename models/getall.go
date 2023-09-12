package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Presensi struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NamaMk         string             `json:"nama_mk" bson:"nama_mk"`
	Tanggal        string             `json:"tanggal" bson:"tanggal"`
	Checkin        string             `json:"checkin" bson:"checkin"`
	Nama_matkul    string             `json:"Nama_matkul,omitempty" bson:"Nama_matkul,omitempty"`
	SKS            string             `json:"SKS,omitempty" bson:"SKS,omitempty"`
	Dosen_pengampu string             `json:"Dosen_pengampu,omitempty" bson:"Dosen_pengampu,omitempty"`
	Email          string             `json:"Email,omitempty" bson:"Email,omitempty"`
	Nama_mhs       Mahasiswa          `json:"Nama_mhs" bson:"Nama_mhs" validate:"required"`
	NPM            string             `json:"NPM" bson:"No NPM" validate:"required"`
	Jurusan        string             `json:"Jurusan" bson:"Jurusan" validate:"required"`
	NPM_ms         string             `json:"NPM_ms,omitempty" validate:"required"`
	Presensi       Absensi            `json:"Presensi,omitempty"`
	Nilai_akhir    string             `json:"Nilai_akhir,omitempty" validate:"required"`
	Grade          string             `json:"Grade,omitempty" validate:"required"`
	Tahun_ajaran   string             `json:"tahun_ajaran,omitempty" validate:"required"`
	Nama_ortu      string             `json:"Nama_ortu,omitempty" validate:"required"`
	Phone_number   string             `json:"Phone_number,omitempty" validate:"required"`
}
type AbsensiNilaiCombined struct {
	AbsensiData Absensi
	NilaiData   Nilai
}
