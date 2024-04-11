package post

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      uint
	Title   string `json:"title"`
	Content string `json:"content"`
	Regdate time.Time
}

func connectDB() (*sql.DB, error) {
	username := "root"
	password := "master"
	dbname := "gin1"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbname))
	if err != nil {
		return nil, err
	}

	// 데이터베이스 연결 테스트
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")
	return db, nil
}

func Posts() []Post {
	db, err := connectDB()

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM post")

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var postList []Post
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

func WritePost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connectDB()
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
