package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/thanhpv3380/api/errors"
	"github.com/thanhpv3380/api/logger"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		logger.Error("Error casting validation error", err)
		return errors.NewError("", "")
	}

	details := make(map[string]string)
	for _, fe := range validationErrors {
		field := fe.Field()
		details[field] = fmt.Sprintf("must be %s", fe.Tag())
	}

	return errors.NewValidationError(details)
}
