package handlers

import (
	"encoding/json"
	"go-auth/internal/config"
	"go-auth/internal/models"
	"go-auth/internal/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const StatusBooked = "booked"

type AppointmentInput struct {
	DoctorID        string    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	var input AppointmentInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendResponse(w, false, "Invalid JSON", nil)
		return
	}

	patientID, ok := r.Context().Value("userID").(string)
	if !ok {
		utils.SendResponse(w, false, "Unauthorized", nil)
		return
	}

	patientObjID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		utils.SendResponse(w, false, "Invalid patient ID", nil)
		return
	}

	doctorObjID, err := primitive.ObjectIDFromHex(input.DoctorID)
	if err != nil {
		utils.SendResponse(w, false, "Invalid doctor ID", nil)
		return
	}

	collection := config.GetCollection(config.Client, "appointments")

	now := time.Now()

	appointment := models.Appointment{
		PatientID:       patientObjID,
		DoctorID:        doctorObjID,
		AppointmentDate: input.AppointmentDate,
		Status:          StatusBooked,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	result, err := collection.InsertOne(ctx, appointment)
	if err != nil {
		utils.SendResponse(w, false, "Failed to create appointment", nil)
		return
	}

	utils.SendResponse(w, true, "Appointment created successfully", result.InsertedID)
}

func GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collection := config.GetCollection(config.Client, "appointments")

	filter := bson.M{}

	// Filtering
	if doctorID := r.URL.Query().Get("doctor_id"); doctorID != "" {
		objID, err := primitive.ObjectIDFromHex(doctorID)
		if err != nil {
			utils.SendResponse(w, false, "Invalid doctor_id", nil)
			return
		}
		filter["doctor_id"] = objID
	}

	if patientID := r.URL.Query().Get("patient_id"); patientID != "" {
		objID, err := primitive.ObjectIDFromHex(patientID)
		if err != nil {
			utils.SendResponse(w, false, "Invalid patient_id", nil)
			return
		}
		filter["patient_id"] = objID
	}

	if date := r.URL.Query().Get("date"); date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			utils.SendResponse(w, false, "Invalid date format (use YYYY-MM-DD)", nil)
			return
		}

		filter["appointment_date"] = bson.M{
			"$gte": parsedDate,
		}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		utils.SendResponse(w, false, "Failed to fetch appointments", nil)
		return
	}
	defer cursor.Close(ctx)

	var appointments []models.Appointment

	for cursor.Next(ctx) {
		var appointment models.Appointment
		if err := cursor.Decode(&appointment); err != nil {
			utils.SendResponse(w, false, "Failed to decode appointment", nil)
			return
		}
		appointments = append(appointments, appointment)
	}

	if err := cursor.Err(); err != nil {
		utils.SendResponse(w, false, "Cursor error", nil)
		return
	}

	utils.SendResponse(w, true, "Appointments fetched successfully", appointments)
}

func GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idHex, ok := params["id"]
	if !ok {
		utils.SendResponse(w, false, "Missing ID parameter", nil)
		return
	}

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		utils.SendResponse(w, false, "Invalid ID", nil)
		return
	}

	collection := config.GetCollection(config.Client, "appointments")

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: id},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "doctors"},
				{Key: "localField", Value: "doctor_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "doctor"},
			}},
		},
		{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$doctor"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		utils.SendResponse(w, false, "Failed to fetch appointment", nil)
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		utils.SendResponse(w, false, "Failed to decode appointment", nil)
		return
	}

	if len(results) == 0 {
		utils.SendResponse(w, false, "Appointment not found", nil)
		return
	}

	utils.SendResponse(w, true, "Appointment fetched successfully", results[0])
}
