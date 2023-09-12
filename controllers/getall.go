package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	GetallCollection     *mongo.Collection = configs.GetCollection(configs.DB, "Abensi")
	mahasiswaCollection  *mongo.Collection = configs.GetCollection(configs.DB, "mahasiswa")
	matakuliahCollection *mongo.Collection = configs.GetCollection(configs.DB, "matakuliah")
	NilaiCollection      *mongo.Collection = configs.GetCollection(configs.DB, "nilai")
	OrangtuaCollection   *mongo.Collection = configs.GetCollection(configs.DB, "orangtua")
)
var Presensi *mongo.Collection = configs.GetCollection(configs.DB, "matakuliah.Abensi")

func getMongoCollections() (
	*mongo.Collection,
	*mongo.Collection,
	*mongo.Collection,
	*mongo.Collection,
	*mongo.Collection,
	error,
) {
	db := configs.DB
	if db == nil {
		// Inisialisasi koneksi ke database MongoDB di sini
	}

	GetallCollection := configs.GetCollection(db, "Abensi")
	mahasiswaCollection := configs.GetCollection(db, "mahasiswa")
	matakuliahCollection := configs.GetCollection(db, "matakuliah")
	NilaiCollection := configs.GetCollection(db, "nilai")
	OrangtuaCollection := configs.GetCollection(db, "orangtua")

	return GetallCollection, mahasiswaCollection, matakuliahCollection, NilaiCollection, OrangtuaCollection, nil
}
func findAllCollections(ctx context.Context, collections ...*mongo.Collection) ([]models.Presensi, error) {
	var allPresensi []models.Presensi

	for _, coll := range collections {
		results, err := coll.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		defer results.Close(ctx)

		var presensi []models.Presensi
		for results.Next(ctx) {
			var presensiDoc models.Presensi
			if err := results.Decode(&presensiDoc); err != nil {
				return nil, err
			}
			presensi = append(presensi, presensiDoc)
		}

		allPresensi = append(allPresensi, presensi...)
	}

	return allPresensi, nil
}

func GetAllDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var (
			absensis    []models.Absensi
			mahasiswas  []models.Mahasiswa
			matakuliahs []models.Matakuliah
			nilais      []models.Nilai
			orangtuas   []models.OrangTua
		)

		// Query data dari koleksi Abensi
		absensiResults, err := GetallCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer absensiResults.Close(ctx)
		for absensiResults.Next(ctx) {
			var absensi models.Absensi
			if err := absensiResults.Decode(&absensi); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			absensis = append(absensis, absensi)
		}

		// Query data dari koleksi Mahasiswa
		// Query data dari koleksi Mahasiswa
		mahasiswaResults, err := mahasiswaCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer mahasiswaResults.Close(ctx)
		for mahasiswaResults.Next(ctx) {
			var mahasiswa models.Mahasiswa
			if err := mahasiswaResults.Decode(&mahasiswa); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			mahasiswas = append(mahasiswas, mahasiswa)
		}

		matakuliahResults, err := MatakuliahCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer matakuliahResults.Close(ctx)
		for matakuliahResults.Next(ctx) {
			var matakuliah models.Matakuliah
			if err := matakuliahResults.Decode(&matakuliah); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			matakuliahs = append(matakuliahs, matakuliah)
		}

		nilaiResults, err := nilaiCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer nilaiResults.Close(ctx)
		for nilaiResults.Next(ctx) {
			var nilai models.Nilai
			if err := nilaiResults.Decode(&nilai); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			nilais = append(nilais, nilai)
		}
		orangtuaResults, err := orangTuaCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer orangtuaResults.Close(ctx)
		for orangtuaResults.Next(ctx) {
			var orangtua models.OrangTua
			if err := orangtuaResults.Decode(&orangtua); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			orangtuas = append(orangtuas, orangtua)
		}
		// collections := []*mongo.Collection{GetallCollection, mahasiswaCollection, matakuliahCollection, NilaiCollection, OrangtuaCollection}
		// var presensis []models.Presensi

		// for _, coll := range collections {
		// 	results, err := coll.Find(ctx, bson.M{})
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 		return
		// 	}
		// 	defer results.Close(ctx)

		// 	for results.Next(ctx) {
		// 		var presensi models.Presensi
		// 		if err := results.Decode(&presensi); err != nil {
		// 			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 			return
		// 		}
		// 		presensis = append(presensis, presensi)
		// 	}
		// }

		// Sekarang, Anda memiliki semua hasil dari semua koleksi dalam slice "presensis"
		// Anda dapat melanjutkan dengan pemrosesan atau respons sesuai kebutuhan Anda

		// ... (lakukan hal yang sama untuk koleksi lainnya)
		// Menggabungkan hasil dari semua koleksi
		allData := map[string]interface{}{
			"absensi":    absensis,
			"mahasiswa":  mahasiswas,
			"matakuliah": matakuliahs,
			"nilai":      nilais,
			"orangtua":   orangtuas,
			// "presensi":   presensis,
		}

		c.JSON(http.StatusOK, responses.AbsensiResponse{Status: http.StatusOK, Message: "success", Data: allData})
	}
}
