package main

import (
	"github.com/ShubhamTatvamasi/GoBackendApp/database"
)

func main() {
	db, err := database.Connect().DB()
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer db.Close()
}
