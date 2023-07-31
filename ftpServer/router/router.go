package router

import (
	"ftpServer/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter1(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/ftpServer")
	apiRouter.POST("/upload/", controller.UploadFile)

}
