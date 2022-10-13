package main

import (
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/router"
)

func main() {
	database.StartDB()

	var PORT = ":8000"

	router.StartServer().Run(PORT)
}
