package main

import (
	"fmt"
	"restapi/initializers"
	"restapi/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Task{})
	fmt.Println("Migrate Success")
}
