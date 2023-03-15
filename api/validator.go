package api

import (
	"simplebank/util"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var validateCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency := fl.Field().String()
	return util.IsSupportedCurrency(currency)
}
