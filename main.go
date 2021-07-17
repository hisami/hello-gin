package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func main() {
	// DB接続
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open database")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()

	// ルーティング
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!!",
		})
	})

	router.Run()
}
