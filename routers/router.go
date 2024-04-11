package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	post "github.com/seanlee0923/first-gin/models/post"
	site "github.com/seanlee0923/first-gin/models/site"
)

func DefaultRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/posts", posts)

	router.POST("/posts", writePost)

	router.GET("/sites", sites)

	router.POST("/sites", writeSite)

	return router
}

func posts(c *gin.Context) {
	result := post.Posts()

	c.JSON(http.StatusOK, gin.H{"message": "조회 성공", "resultArray": result})
}

func writePost(c *gin.Context) {
	post.WritePost(c)
}

func sites(c *gin.Context) {
	result := site.Sites()

	c.JSON(http.StatusOK, gin.H{"message": "조회 성공", "resultArray": result})
}

func writeSite(c *gin.Context) {
	site.WriteSite(c)
}
