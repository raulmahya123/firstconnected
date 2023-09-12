package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Matakuliah struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama_matkul    string             `json:"Nama_matkul,omitempty" bson:"Nama_matkul,omitempty"`
	SKS            string             `json:"SKS,omitempty" bson:"SKS,omitempty"`
	Dosen_pengampu string             `json:"Dosen_pengampu,omitempty" bson:"Dosen_pengampu,omitempty"`
	Email          string             `json:"Email,omitempty" bson:"Email,omitempty"`
}
