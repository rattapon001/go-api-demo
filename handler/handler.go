package handler

import (
	"demo1/database/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users := []model.User{}

	h.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := model.User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	fmt.Println("user.id =>>> ", user.Id)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   nil,
		})
	} else {
		data := []model.User{}
		data = append(data, user)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})
	}

}

func (h *UserHandler) SaveUser(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
