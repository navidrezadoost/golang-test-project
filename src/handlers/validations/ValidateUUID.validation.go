package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsUUID(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	return uuidRegex.MatchString(value)
}
