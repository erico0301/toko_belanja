package main

import (
	"toko_belanja/database"
	"toko_belanja/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := "8080"
	r.Run(":" + port)
}
