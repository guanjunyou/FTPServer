package impl

import (
	"errors"
	"ftpServer/config"
	"ftpServer/service"
	"ftpServer/utils"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os/exec"
	"path"
	"strconv"
)

type FileServiceImpl struct {
	service.FileService
}

func (fileService FileServiceImpl) SaveFile(c *gin.Context, file *multipart.FileHeader, fileType int64) (string, string, error) {

	var savePath string
	if fileType == 1 {
		savePath = path.Join(config.CommonFilePath, config.VideoDir)
	} else {
		savePath = path.Join(config.CommonFilePath, config.PhotoDir)
	}

	url, err := utils.SaveFileToFileSystem(c, file, savePath)
	if err != nil {
		return "", "", err
	}
	if fileType == 1 {
		log.Println(path.Join(savePath, url))
		coverUrl, err1 := SaveCoverFile(path.Join(savePath, url))
		if err1 != nil {
			return url, "", err1
		}
		return url, coverUrl, nil
	} else {
		return url, "", nil
	}

}

func SaveCoverFile(videoPath string) (string, error) {
	var coverPath string
	nextID := utils.NewSnowflake().NextID()
	coverName := config.CommonCoverName + strconv.FormatInt(nextID, 10) + ".jpg"
	coverPath = path.Join(config.CommonFilePath, config.PhotoDir, coverName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:01", "-vframes", "1", coverPath)
	err := cmd.Run()
	if err != nil {
		log.Println("截图失败", err)
		return "", errors.New("截图失败")
	}
	return coverName, nil
}
