package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Site struct {
	Id             uint
	SiteName       string `json:"siteName"`
	SiteInfo       string `json:"siteInfo"`
	SiteIP         string `json:"siteIp"`
	SitePort       uint   `json:"sitePort"`
	SiteUsername   string `json:"siteUsername"`
	SitePassword   string `json:"sitePassword"`
	SiteOSName     string `json:"siteOSName"`
	SiteOSPassword string `json:"siteOSPassword"`
	IsBasic        bool   `json:"isBasic"`
	CreatedAt      time.Time
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

func Sites() []Site {
	var sites []Site

	db, err := connectDB()
	if err != nil {
		panic("hi")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM site")
	if err != nil {
		panic("hi")
	}
	defer rows.Close()
	for rows.Next() {
		var site Site
		rows.Scan(
			&site.Id, &site.SiteName,
			&site.SiteInfo,
			&site.SiteIP, &site.SitePort,
			&site.SiteUsername, &site.SitePassword,
			&site.SiteOSName, &site.SiteOSPassword,
			&site.IsBasic, &site.CreatedAt)
		sites = append(sites, site)
	}
	return sites
}

func WriteSite(c *gin.Context) {
	var site Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connectDB()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO site (sitename, siteinfo, siteip, siteport, siteusername, sitepassword, siteosname, siteospassword, isbasic) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		site.SiteName, site.SiteInfo,
		site.SiteIP, site.SitePort,
		site.SiteUsername, site.SitePassword,
		site.SiteOSName, site.SiteOSPassword,
		site.IsBasic)
	if err != nil {
		panic(err.Error())
	}
}
