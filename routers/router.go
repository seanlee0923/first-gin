package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanlee0923/first-gin/models"
)

func DefaultRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/posts", posts)

	router.POST("/posts", writePost)

	return router
}

func posts(c *gin.Context) {
	result := models.Posts()

	c.JSON(http.StatusOK, gin.H{"message": "조회 성공", "resultArray": result})
}

func writePost(c *gin.Context) {
	models.WritePost(c)
}
