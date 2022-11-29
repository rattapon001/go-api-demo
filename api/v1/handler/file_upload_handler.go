package handler

import (
	"context"
	"demo1/internal/entity"
	minioClient "demo1/internal/minio_client"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type FileHandler struct {
	DB *gorm.DB
}

func (h *FileHandler) UploadFile(c *gin.Context) {

	var contentType string
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET")

	file, _ := c.FormFile("file")
	contentType = file.Header.Get("Content-Type")
	f, err := file.Open()

	minioUpload := minioClient.MinioClientSetup()
	current_time := time.Now()

	fileName := current_time.Format(time.RFC3339Nano) + "-" + file.Filename

	info, err := minioUpload.PutObject(ctx, bucketName, "demo1/"+fileName, f, file.Size, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data":   nil,
		})
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, info.Size)

	mediaObject := entity.MediaObject{}

	mediaObject.DirectoryName = "demo1"
	mediaObject.FilePath = fileName

	if err := h.DB.Save(&mediaObject).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resData := []entity.MediaObject{}
	resData = append(resData, mediaObject)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   resData,
	})
}
