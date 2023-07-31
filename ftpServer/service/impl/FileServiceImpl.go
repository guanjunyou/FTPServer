package impl

import (
	"ftpServer/config"
	"ftpServer/service"
	"ftpServer/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
)

type FileServiceImpl struct {
	service.FileService
}

func (fileService FileServiceImpl) SaveFile(c *gin.Context, file *multipart.FileHeader, fileType int64) (string, error) {

	var savePath string
	if fileType == 1 {
		savePath = path.Join(config.CommonFilePath, config.VideoDir)
	} else {
		savePath = path.Join(config.CommonFilePath, config.PhotoDir)
	}

	url, err := utils.SaveFileToFileSystem(c, file, savePath)
	if err != nil {
		return "", err
	}
	return url, nil
}
