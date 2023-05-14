package main

import (
	"btpn-fullstack/database"
	"btpn-fullstack/models"
	"btpn-fullstack/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
