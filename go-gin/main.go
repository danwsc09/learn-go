package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	db := GetDbConnection()
	defer db.Close()
	/*
		router.GET("/albums", GetAlbums)
		router.GET("/albums/:id", GetAlbumByID)
		router.POST("/albums", PostAlbums)
	*/

	router.GET("/users", GetUsers)
	router.GET("/users/:minAge", GetUsersFilterByAge)
	router.POST("/users", PostUsers)
	router.Run("localhost:8080")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
