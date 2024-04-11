package post

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/seanlee0923/first-gin/models"
)

type Post struct {
	Id      uint
	Title   string `json:"title"`
	Content string `json:"content"`
	Regdate time.Time
}

func Posts() []Post {
	var postList []Post

	db, err := models.ConnectDB()

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM post")

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.Regdate)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Id: ", post.Id)
		log.Println("Id: ", post.Title)
		log.Println("Id: ", post.Content)
		log.Println("Id: ", post.Regdate)
		postList = append(postList, post)
	}
	return postList
}

func PostBy(id uint) Post {
	var post Post

	db, err := models.ConnectDB()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	row, err := db.Query("SELECT * FROM post WHERE ID = ?", id)

	if err != nil {
		panic(err.Error())
	}
	row.Scan(&post.Id, &post.Title, &post.Content, &post.Regdate)

	return post
}

func WritePost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := models.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO post (title, content) VALUES (?, ?)", post.Title, post.Content)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post Created Successfully"})
}
