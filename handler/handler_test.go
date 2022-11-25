package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestUserHandler_GetUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				DB: tt.fields.DB,
			}
			h.GetUser(tt.args.c)
		})
	}
}
