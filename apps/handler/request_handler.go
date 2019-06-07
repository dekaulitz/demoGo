package handler

import (
	"demoGo/apps/handler/exception"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strings"
)

/**
validating Object struct with govalidator rule
*/
func ValidateRequest(req interface{}) *exception.ErrorException {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return exception.NewException(exception.VALIDATION_FAIL).Throw(err.Error())
	}
	return nil
}

/**
create raw URL from current URL binding with gin.Params and replacing with the key
*/
func GetRawURL(url string, params gin.Params) string {
	for _, element := range params {
		url = strings.Replace(url, element.Value, ":"+element.Key, -1)
	}
	return url
}
