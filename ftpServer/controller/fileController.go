package controller

import (
	"ftpServer/models"
	"ftpServer/service/impl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type FileUrlResponse struct {
	models.Response
	Url      string `json:"url"`
	CoverUrl string `json:"cover_url"`
}

func GetFileService() impl.FileServiceImpl {
	var FileService impl.FileServiceImpl
	return FileService
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	fileType := c.PostForm("fileType") // 1 为视频 2 为图片
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	fileTypeInt, _ := strconv.ParseInt(fileType, 10, 64)
	url, coverUrl, err1 := GetFileService().SaveFile(c, file, fileTypeInt)
	if err1 != nil {
		log.Printf(err1.Error())
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "保存失败！",
		})
		return
	}
	c.JSON(http.StatusOK, FileUrlResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "保存成功",
		},
		Url:      url,
		CoverUrl: coverUrl,
	})

}
