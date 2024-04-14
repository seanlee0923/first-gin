package routers

import (
	"net/http"
	"strconv"

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

	router.POST("/posts", writePost)

	router.GET("/posts", posts)

	router.GET("/post", postBy)

	router.POST("/post-update/:id", updatePost)

	router.POST("/post-delete", deletePostBy)

	router.POST("/sites", writeSite)

	router.GET("/sites", sites)

	router.GET("/site", getSiteBy)

	router.POST("/site-update", updateSite)

	router.POST("/site-delete", deleteSiteBy)

	return router
}

func writePost(c *gin.Context) {
	post.WritePost(c)
}

func posts(c *gin.Context) {
	result := post.Posts()

	c.JSON(http.StatusOK, gin.H{"message": "조회 성공", "resultArray": result})
}

func postBy(c *gin.Context) {
	result := post.PostBy(2)

	c.JSON(http.StatusOK, gin.H{"message": "단건 조회 성공", "result": result})
}

func updatePost(c *gin.Context) {
	var changedPost post.Post
	postId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := c.ShouldBindJSON(&changedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := post.UpdatePost(changedPost, postId)
	c.JSON(http.StatusOK, gin.H{"message": "업데이트 해보자", "result": result})
}

func deletePostBy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "삭제도 해보자"})
}

func writeSite(c *gin.Context) {
	site.WriteSite(c)
}

func sites(c *gin.Context) {
	result := site.Sites()

	c.JSON(http.StatusOK, gin.H{"message": "조회 성공", "resultArray": result})
}

func getSiteBy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "단건조회 테스트"})
}

func updateSite(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "업데이트 테스트"})
}

func deleteSiteBy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "삭제 테스트"})
}
