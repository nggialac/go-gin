package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lacnguyen/go-gin/entity"
	"github.com/lacnguyen/go-gin/service"
	"github.com/lacnguyen/go-gin/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
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

func (vc *videoController) Update(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	vc.videoService.Update(video)
	return nil
}

func (vc *videoController) Delete(c *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	vc.videoService.Delete(video)
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

func (vc *videoController) FindAll() []entity.Video {
	return vc.videoService.FindAll()
}
