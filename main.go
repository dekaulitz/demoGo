package main

import (
	"bitbucket.com/LippoDigitalOVO/ovo-auth/lib/uuid"
	"bytes"
	"demoGo/apps/controllers"
	"demoGo/apps/helper"
	"demoGo/apps/routes"
	"demoGo/configuration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	config = configuration.GetConfiguration()
)

func main() {
	router := setupRouter()
	router.Run(config.Host.Host + ":" + config.Host.Port)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		blw := &helper.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		var logBody helper.LogRequest
		t := time.Now()
		logBody.RequestId = uuid.GenerateUUID()
		logBody.Url = c.Request.RequestURI
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
		logBody.Body = helper.SpaceMap(string(requestBodyBytes))
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
	})
	// Cors handler
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "token", "transaction_type_id", "merchant_id", "merchant_invoice", "title", "subcategory_title", "article_title"},
		AllowAllOrigins: true,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "message": "routing not found"})
	})
	router.GET("/health", controllers.HealthCheck)
	routes.SetRouterv1(router)
	return router
}
