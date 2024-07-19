package api

import (
	"github.com/daisydbbb/Simple-Bank/db/util"
	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currentcy, ok :=fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currentcy)
	}
	return false
}