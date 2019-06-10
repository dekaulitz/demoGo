package middleware

/**
this go file aim for grouping middleware handler for easy maintain
you can add middleware handler on this file and injecting to the routes
*/
import (
	"bytes"
	"demoGo/apps/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"time"
)

/**
this is global filter will log the request and also the response
the log is already with the struct if you want to pass to the other thirdpary like new relic or datadog
you can create handler for send the log to them
*/
func GlobalHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Request.Header.Add("request-id", uuid.New().String())
	blw := &handler.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	var logBody handler.LogRequest
	t := time.Now()
	logBody.RequestId = c.Request.Header.Get("request-id")
	logBody.Url = c.Request.RequestURI
	var requestBodyBytes []byte
	if c.Request.Body != nil {
		requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
	logBody.Body = handler.SpaceMap(string(requestBodyBytes))
	if logBody.Body == "" {
		logBody.Body = "nil"
	}
	logBody.Method = c.Request.Method
	c.Next()
	logBody.Header = c.Writer.Header()
	logBody.Latency = time.Since(t)
	logBody.HttpCode = c.Writer.Status()
	logBody.ResponseBody = blw.Body.String()
	go func() {
		log.Printf("[INFO]%v REQUESTID : %s, HEADER: %v, URL : %v, METHOD : %v,"+
			"REQUESTBODY: %v, LATENCY: %v,STATUS: %v,RESPONSEBODY : %v \n", time.Now().Format(time.RFC3339),
			c.Request.Header.Get("request-id"), logBody.Header, logBody.Url, logBody.Method, logBody.Body,
			logBody.Latency, logBody.HttpCode, logBody.ResponseBody,
		)
	}()
}

/**
this is for cors handler if the request consuming from web
*/
func CorsHandler(router *gin.Engine) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "token", "transaction_type_id", "merchant_id", "merchant_invoice", "title", "subcategory_title", "article_title"},
		AllowAllOrigins: true,
	})
}

//func HandleErrors(c *gin.Context) {
//	c.Next() // execute all the handlers
//
//	// at this point, all the handlers finished. Let's read the errors!
//	if len(c.Errors) == 0 {
//		return
//	}
//fmt.Println("error")
//	c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
//}
