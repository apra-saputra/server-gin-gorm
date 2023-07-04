package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"restapi/initializers"
)

type SeedData struct {
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// Baca file JSON
	data, errJson := ioutil.ReadFile("seed/data.json")
	if errJson != nil {
		log.Fatal("Error Read JSON", errJson)
	}

	// Parse data JSON
	var seedData SeedData
	errParse := json.Unmarshal(data, &seedData)
	if errParse != nil {
		log.Fatal("Error Parse JSON", errParse)
	}

	// Insert data ke database
	for _, user := range seedData.Users {
		err := initializers.DB.Create(&user).Error
		if err != nil {
			log.Fatal("Error Write User", user, "error : ", err)
		}
	}

	fmt.Println("Seeder User Success")
}
