package handler

import (
	"demoGo/apps/handler/exception"
	"github.com/asaskevich/govalidator"
)

func ValidateRequest(req interface{}) *exception.ErrorException {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return exception.NewException(exception.VALIDATION_FAIL).Throw(err.Error())
	}
	return nil
}
