package main

import (
	"fmt"
	"restapi/initializers"
	"restapi/services/model"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
	initializers.DB.AutoMigrate(&model.Task{})
	fmt.Println("Migrate Success")
}
