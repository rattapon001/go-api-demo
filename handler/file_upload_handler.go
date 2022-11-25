package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"demo1/database/model"
	minioClient "demo1/minio_client"

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
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, info.Size)

	mediaObject := model.MediaObject{}

	mediaObject.DirectoryName = "demo1"
	mediaObject.FilePath = fileName

	if err := h.DB.Save(&mediaObject).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resData := []model.MediaObject{}
	resData = append(resData, mediaObject)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   resData,
	})
}
