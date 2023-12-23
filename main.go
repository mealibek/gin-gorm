package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mealibek/gin-gorm/controllers"
	"github.com/mealibek/gin-gorm/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsList)
	r.GET("/posts/:id", controllers.PostsRetrieve)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run(fmt.Sprintf("localhost:%v", os.Getenv("PORT")))
}
