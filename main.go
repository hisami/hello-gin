package main

import (
	"hello-gin/domain/model"
	"hello-gin/infra"
	"hello-gin/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DB接続
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open database")
	}
	db.AutoMigrate(&model.Todo{})
	defer db.Close()

	// DI
	todoRepository := infra.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	// ルーティング
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		todos, _ := todoUsecase.FindAll()
		c.JSON(http.StatusOK, todos) // 構造体の配列
	})

	router.POST("/", func(c *gin.Context) {
		type TodoRequest struct {
			Title  string `json:"title"`
			Status string `json:"status"`
		}
		var todoRequest TodoRequest
		c.BindJSON(&todoRequest)
		todoUsecase.Create(todoRequest.Title, todoRequest.Status)
		c.JSON(http.StatusOK, gin.H{
			"status": "success!!",
		})
	})

	router.Run()
}
