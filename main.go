package main

import (
	"restapi/initializers"
	"restapi/router"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.ConnectRedis()
}

func main() {
	router := router.SetupRouter()

	router.Run() // listen and serve on 0.0.0.0:3001
}
