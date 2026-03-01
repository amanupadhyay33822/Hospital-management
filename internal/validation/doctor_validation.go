package validation

type CreateDoctor struct {
	Name       string `json:"name" validate:"required,min=3"`
	Email      string `json:"email" validate:"required,email"`
	Speciality string `json:"speciality" validate:"required,min=5,alpha"`
}

func ValidateDoctor(input CreateDoctor) error {
	return validate.Struct(input)
}
