package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Absensi struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Nama_mk string             `json:"Nama_mk,omitempty"`
	Tanggal string             `json:"Tanggal,omitempty"`
	Checkin string             `json:"Checkin,omitempty"`
}
