package main

import (
	"ftpServer/config"
	"ftpServer/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ReadConfig()
	r := gin.Default()
	router.InitRouter1(r)
	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8080")
}
