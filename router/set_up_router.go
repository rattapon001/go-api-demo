package router

import (
	"bytes"
	"demo1/database"
	"demo1/handler"
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
	fmt.Println("Response body: " + w.body.String())
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logResponseBody)
	db := database.Initialize()
	UserRouter(r, db)
	r.GET("/ping")

	fileHandler := handler.FileHandler{}
	fileHandler.DB = db
	r.POST("/upload", fileHandler.UploadFile)
	// r.PUT("/customers/:id", h.UpdateCustomer)
	// r.DELETE("/customers/:id", h.DeleteCustomer)

	return r
}
