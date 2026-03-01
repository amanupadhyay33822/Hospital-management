package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PatientID       primitive.ObjectID `json:"patient_id" bson:"patient_id"`
	DoctorID        primitive.ObjectID `json:"doctor_id" bson:"doctor_id"`
	AppointmentDate time.Time          `json:"appointment_date" bson:"appointment_date"`
	Status          string             `json:"status" bson:"status"` // booked, cancelled, completed
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
