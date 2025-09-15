package contact

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateContactDTO struct {
	FullName string `json:"fullname" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Subject  string `json:"subject" validate:"required"`
	Message  string `json:"message" validate:"required"`
}

func ValidateDTO(dto interface{}) error {
	return validate.Struct(dto)
}
