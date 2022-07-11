package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	db := GetDbConnection()
	defer db.Close()
	/*
		router.GET("/albums", GetAlbums)
		router.GET("/albums/:id", GetAlbumByID)
		router.POST("/albums", PostAlbums)
	*/

	router.GET("/users", GetUsers)
	router.Run("localhost:8080")
}
