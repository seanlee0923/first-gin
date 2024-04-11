package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/seanlee0923/first-gin/models"
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

func Sites() []Site {
	var sites []Site

	db, err := models.ConnectDB()
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
	fmt.Println(site.SiteName)
	db, err := models.ConnectDB()

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
