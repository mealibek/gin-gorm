package main

import (
	"github.com/mealibek/gin-gorm/initializers"
	"github.com/mealibek/gin-gorm/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
