package handlers

import (
	"context"
	"encoding/json"
	"go-auth/internal/config"
	"go-auth/internal/models"
	"go-auth/internal/utils"
	"go-auth/internal/validation"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var input validation.CreateDoctor

	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.SendResponse(w, false, "Invalid json", nil)
		return
	}

	// validate input
	err = validation.ValidateDoctor(input)
	if err != nil {
		utils.SendResponse(w, false, validation.FormatValidationError(err), nil)
		return
	}
	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "doctors")

	//check duplicate email
	var existDoctor models.Doctor
	err = collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&existDoctor)
	if err == nil {
		utils.SendResponse(w, false, "Email already exists", nil)
		return
	}
	if err != mongo.ErrNoDocuments {
		utils.SendResponse(w, false, "Database error", nil)
		return
	}
	var newDoctor models.Doctor
	newDoctor = models.Doctor{
		ID:         primitive.NewObjectID(),
		Name:       input.Name,
		Email:      input.Email,
		Speciality: input.Speciality,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), newDoctor)
	if err != nil {
		utils.SendResponse(w, false, "Failed to create doctor", nil)
		return
	}

	utils.SendResponse(w, true, "Doctor created successfully", newDoctor)

}

func GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	collection := config.GetCollection(config.Client, "doctors")

	// Default pagination
	page := 1
	limit := 10

	// Get query params directly
	p := r.URL.Query().Get("page")
	l := r.URL.Query().Get("limit")

	if p != "" {
		parsedPage, err := strconv.Atoi(p)
		if err != nil || parsedPage <= 0 {
			utils.SendResponse(w, false, "Invalid page number", nil)
			return
		}
		page = parsedPage
	}

	if l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil || parsedLimit <= 0 {
			utils.SendResponse(w, false, "Invalid limit number", nil)
			return
		}
		limit = parsedLimit
	}

	skip := (page - 1) * limit

	findOptions := options.Find().
		SetSort(bson.M{"created_at": -1}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		utils.SendResponse(w, false, "Failed to fetch doctors", nil)
		return
	}
	defer cursor.Close(ctx)

	var doctors []models.Doctor

	for cursor.Next(ctx) {
		var doctor models.Doctor
		if err := cursor.Decode(&doctor); err != nil {
			utils.SendResponse(w, false, "Failed to decode doctor", nil)
			return
		}
		doctors = append(doctors, doctor)
	}

	if err := cursor.Err(); err != nil {
		utils.SendResponse(w, false, "Cursor error", nil)
		return
	}

	utils.SendResponse(w, true, "Doctors fetched successfully", doctors)
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {

	var params map[string]string
	params = mux.Vars(r)
	var id primitive.ObjectID
	var err error
	id, err = primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendResponse(w, false, "Invalid ID", nil)
		return
	}
	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "doctors")

	var doctor models.Doctor
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&doctor)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendResponse(w, false, "Doctor not found", nil)
			return
		}
		utils.SendResponse(w, false, "Failed to fetch doctor", nil)
		return
	}
	utils.SendResponse(w, true, "Doctor fetched successfully", doctor)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	var input models.Doctor

	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.SendResponse(w, false, "Invalid JSON", nil)
		return
	}
	var params map[string]string
	params = mux.Vars(r)
	var id primitive.ObjectID
	id, err = primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendResponse(w, false, "Invalid ID", nil)
		return
	}
	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "doctors")
	var doctor models.Doctor
	var updateBody bson.M
	updateBody = bson.M{"$set": input}
	err = collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, updateBody).Decode(&doctor)
	if err != nil {
		utils.SendResponse(w, false, "Failed to update doctor", nil)
		return
	}
	utils.SendResponse(w, true, "Doctor updated successfully", doctor)
}

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	params = mux.Vars(r)
	var id primitive.ObjectID
	var err error
	id, err = primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendResponse(w, false, "Invalid ID", nil)
		return
	}
	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "doctors")
	var deletedDoctor models.Doctor
	err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&deletedDoctor)
	if err != nil {
		utils.SendResponse(w, false, "Failed to delete doctor", nil)
		return
	}
	utils.SendResponse(w, true, "Doctor deleted successfully", deletedDoctor)
}
