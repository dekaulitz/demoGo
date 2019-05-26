package helper

import (
	"demoGo/apps/vmodel"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOk(body interface{}, c *gin.Context, httpCode int) {
	var response vmodel.ResponseModel
	response.Body = body
	ress, err := json.Marshal(response)
	if err != nil {
		ResponseError(err, c, http.StatusBadRequest)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(httpCode)
	c.Writer.Write(ress)
	return
}

func ResponseError(body interface{}, c *gin.Context, httpcode int) {

}
