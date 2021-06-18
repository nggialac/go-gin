package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lacnguyen/go-gin/entity"
	"github.com/lacnguyen/go-gin/service"
	"github.com/lacnguyen/go-gin/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
}

type videoController struct {
	videoService service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		videoService: service,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.videoService.FindAll()
}

func (vc *videoController) Save(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	vc.videoService.Save(video)
	return nil
}

func (ctl *videoController) ShowAll(c *gin.Context) {
	videos := ctl.videoService.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
