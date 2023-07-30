package service

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FileService interface {
	SaveFile(c *gin.Context, file *multipart.FileHeader, fileType int64) (string, error)
}
