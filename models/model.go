package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	id      uint
	Title   string
	Content string
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

func Posts() {
	db, err := connectDB()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM post")

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var title string
		var content string
		var regdate time.Time

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("id: ", id)
		fmt.Println("title: ", title)
		fmt.Println("content: ", content)
		fmt.Println("regdate: ", regdate)
	}
}
