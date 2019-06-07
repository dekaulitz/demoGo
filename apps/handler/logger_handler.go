package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"unicode"
)

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
