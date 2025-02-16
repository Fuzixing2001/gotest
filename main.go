package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "你gin")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	//默认在0.0.0.0：8080端口启动服务，可以指定端口
	//r.Run(":8081") //启动一个web服务
}
