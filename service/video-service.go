package service

import (
	"github.com/rsengaravua/go-clean/models"
)

// creating interface for usuable functions
type VideoService interface {
	Save(models.Video) models.Video
	GetAll() []models.Video
}

// Structure for VideoService Interface
type videoService struct {
	videos []models.Video
}

// constructor function
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(videos models.Video) models.Video {
	service.videos = append(service.videos, videos)
	return videos
}

func (service *videoService) GetAll() []models.Video {
	return service.videos
}
