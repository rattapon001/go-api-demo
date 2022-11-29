package router

import (
	"bytes"
	"demo1/api/v1/handler"
	"demo1/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func logResponseBody(c *gin.Context) {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	// c.JSON(http.StatusOK, gin.H{
	// 	"status": http.StatusOK,
	// 	"data":   nil,
	// })
	fmt.Println("Response body: " + w.body.String())
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logResponseBody)
	db := config.InitializeDatabase()
	UserRouter(r, db)
	r.GET("/ping")

	fileHandler := handler.FileHandler{}
	fileHandler.DB = db
	r.POST("/upload", fileHandler.UploadFile)

	return r
}
