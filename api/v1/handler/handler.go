package handler

import (
	"demo1/internal/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users := []entity.User{}
	h.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := entity.User{}

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
		data := []entity.User{}
		data = append(data, user)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})
	}

}

func (h *UserHandler) SaveUser(c *gin.Context) {
	user := entity.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	data := []entity.User{}
	data = append(data, user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := entity.User{}
	body := entity.User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Age = body.Age
	user.Email = body.Email

	if err := h.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	data := []entity.User{}
	data = append(data, user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})

}
