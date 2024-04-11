package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
