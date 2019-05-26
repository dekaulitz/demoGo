package helper

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"unicode"
)

//func SetLogResponse(c *gin.Context, body []byte, httpCode int) {
//	var logger libraries.LogResponseModel
//	logger.Header = c.Request.Header
//	logger.RequestId = c.Request.Header.Get("request-id")
//	logger.UriPath = getRawUrl(c.Request.URL.Path, c.Params)
//	logger.Payload = body
//	logger.HttpCode = httpCode
//	libraries.LoggingResponse(logger)
//
//}
//
//func SetLogRequest(c *gin.Context, body []byte) {
//	var logger libraries.LogRequest
//	logger.Header = c.Request.Header
//	logger.RequestId = c.Request.Header.Get("request-id")
//	logger.UriPath = getRawUrl(c.Request.URL.Path, c.Params)
//	logger.Payload = body
//	libraries.LoggingRequest(logger)
//}

func getRawUrl(url string, params gin.Params) string {
	for _, element := range params {
		url = strings.Replace(url, element.Value, ":"+element.Key, -1)
	}
	return url
}

type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}
type LogRequest struct {
	RequestId    string        `json:"request_id"`
	Header       interface{}   `json:"header"`
	Method       string        `json:"method"`
	Url          string        `json:"url"`
	Body         interface{}   `json:"body"`
	HttpCode     int           `json:"http_code"`
	Latency      time.Duration `json:"latency"`
	ResponseBody interface{}   `json:"response_body"`
}

var ()

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

//remove whitespace
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
