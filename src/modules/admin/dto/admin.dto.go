package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateAdminDTO struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

func ValidateDTO(dto interface{}) error {
	return validate.Struct(dto)
}
