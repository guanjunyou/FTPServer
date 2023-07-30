package main

import (
	"ftpServer/config"
	"ftpServer/router"
	"ftpServer/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ReadConfig()

	utils.InitClient()
	r := gin.Default()
	router.InitRouter1(r)
	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8080")
}
