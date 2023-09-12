package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var nilaiCollection *mongo.Collection = configs.GetCollection(configs.DB, "nilai")
var validate_nilai = validator.New()

func CreateNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nilai models.Nilai
		defer cancel()

		// Validate the request body
		if err := c.BindJSON(&nilai); err != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate_nilai.Struct(&nilai); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Check if the "Absensi" is less than 8 and set "Nilai_akhir" to 40 if it is
		if nilai.Presensi < "8" {
			nilai.Nilai_akhir = "E"
			// } else if nilai.Nilai_akhir < "60" {
			// 	c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Nilai harus lebih dari atau sama dengan 60"}})
			// 	return
		}

		newAbsensi := models.Nilai{
			Id:           primitive.NewObjectID(),
			NPM_ms:       nilai.NPM_ms,
			Presensi:     nilai.Presensi,
			Nilai_akhir:  nilai.Nilai_akhir,
			Grade:        nilai.Grade,
			Tahun_ajaran: nilai.Tahun_ajaran,
		}

		result, err := nilaiCollection.InsertOne(ctx, newAbsensi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.NilaiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		NilaiID := c.Param("nialiGetID")
		var nilai models.Nilai
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(NilaiID)

		err := nilaiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&nilai)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": nilai}})
	}
}

func EditNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		nilaiID := c.Param("nialiID")
		var nilai models.Nilai
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(nilaiID)

		//validate the request body
		if err := c.BindJSON(&nilai); err != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_nilai.Struct(&nilai); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"NPM_ms": nilai.NPM_ms, "Presensi": nilai.Presensi, "Nilai_akhir": nilai.Nilai_akhir, "Grade": nilai.Grade, "tahun ajaran": nilai.Tahun_ajaran}
		result, err := nilaiCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Vertebrata details
		var updatedNilai models.Nilai
		if result.MatchedCount == 1 {
			err := nilaiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedNilai)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedNilai}})
	}
}

func DeleteNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		DeleteID := c.Param("nialiID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(DeleteID)

		result, err := nilaiCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.NilaiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "nilai with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "nilai successfully deleted!"}},
		)
	}
}

func GetAllNilais() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var vertebratas []models.Nilai
		defer cancel()

		results, err := nilaiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleVertebrata models.Nilai
			if err = results.Decode(&singleVertebrata); err != nil {
				c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			vertebratas = append(vertebratas, singleVertebrata)
		}

		c.JSON(http.StatusOK,
			responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": vertebratas}},
		)
	}
}

func GetNPM() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		NPM := c.Param("npm")
		var nilai models.Nilai
		defer cancel()

		// Log the received NPM value
		fmt.Printf("Received NPM value: %s\n", NPM)

		objId, err := primitive.ObjectIDFromHex(NPM)
		if err != nil {
			// Log the error for invalid NPM format
			fmt.Printf("Invalid NPM format error: %v\n", err)
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid NPM format"}})
			return
		}

		// Log the ObjectID obtained from the NPM value
		fmt.Printf("ObjectID from NPM: %s\n", objId.Hex())

		err = nilaiCollection.FindOne(ctx, bson.M{"NPM_ms": objId}).Decode(&nilai)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// Log when Nilai is not found
				fmt.Println("Nilai not found in the database")
				c.JSON(http.StatusNotFound, responses.NilaiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Nilai not found"}})
				return
			}
			// Log other internal server errors
			fmt.Printf("Internal server error: %v\n", err)
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Log the successful response
		fmt.Println("Successfully retrieved Nilai from the database")

		c.JSON(http.StatusOK, responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": nilai}})
	}
}
