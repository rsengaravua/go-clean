package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-clean/controller"
	"github.com/rsengaravua/go-clean/service"
)

var (
	VideoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.GetAll())
	})
	server.POST("/video", func(ctx *gin.Context) {
		VideoController.Save(ctx)
	})
	server.Run(":9090")
}
