package model

import (
	"github.com/go-playground/validator/v10"
)

// BlogPost represents a blog post.
type BlogPost struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required"`
}

// Validate validates the BlogPost struct.
func (bp *BlogPost) Validate() error {
	validate := validator.New()

	// Define custom validation functions
	validate.RegisterValidation("titleLength", func(fl validator.FieldLevel) bool {
		title := fl.Field().String()
		return len(title) >= 5 // Adjust as needed
	})

	// Apply the custom validation rules
	validate.RegisterAlias("titleMinLength", "titleLength=5")
	validate.RegisterAlias("titleMaxLength", "max=100")

	if err := validate.Struct(bp); err != nil {
		return err
	}

	return nil
}
