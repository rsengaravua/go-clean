package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-clean/models"
	"github.com/rsengaravua/go-clean/service"
)

type VideoController interface {
	// return all videos
	GetAll() []models.Video
	// Save the video
	Save(ctx *gin.Context) models.Video
}

type controller struct {
	service service.VideoService
}

// GetAll implements VideoController.
func (v controller) GetAll() []models.Video {
	return v.service.GetAll()
}

// Save implements VideoController.
func (v controller) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	v.service.Save(video)
	return video
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
