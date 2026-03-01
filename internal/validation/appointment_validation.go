package validation

type AppointmentRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
	Date  string `json:"date" validate:"required"`
	Time  string `json:"time" validate:"required"`
}

func ValidateAppointmentRequest(appointment *AppointmentRequest) error {
	return validate.Struct(appointment)
}

