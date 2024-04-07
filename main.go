package main

import (
	"github.com/ShubhamTatvamasi/GoBackendApp/database"
)

func main() {
	db := database.Connect()
	closeDB, err := db.DB()
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer closeDB.Close()

	db.AutoMigrate(&database.User{})
}
