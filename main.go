package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lacnguyen/go-gin/controller"
	"github.com/lacnguyen/go-gin/middlewares"
	"github.com/lacnguyen/go-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	//server := gin.Default()
	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok!",
			})
		})

		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(http.StatusOK, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Video Input is valid!!!",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
