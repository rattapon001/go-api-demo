package router

import (
	"demo1/api/v1/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {

	userHandler := handler.UserHandler{}
	userHandler.DB = db

	r.GET("/users", userHandler.GetAllUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.SaveUser)
	r.PATCH("/users/:id", userHandler.UpdateUser)
}
