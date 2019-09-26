package pkg

import "gopkg.in/go-playground/validator.v9"

// RequestValidator decouples validator lib with our system.
type (
	RequestValidator interface {
		ValidateStruct(str interface{}) error
	}

	validatorImpl struct {
		validate *validator.Validate
	}
)

// NewRequestValidator ...
func NewRequestValidator() RequestValidator {
	return &validatorImpl{validator.New()}
}

func (v *validatorImpl) ValidateStruct(str interface{}) error {
	return v.validate.Struct(str)
}
