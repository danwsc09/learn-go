package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user represents data about a record album.
type UserType struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// GetUsers responds with the list of all albums as JSON.
func GetUsers(c *gin.Context) {
	sqlStatement := `
	SELECT * FROM users`
	db := GetDbConnection()
	defer db.Close()
	res, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Error:", err)
		c.Status(500)
	}
	res.RowsAffected()
	fmt.Println("res:", res)
	c.IndentedJSON(http.StatusOK, res)
}

/*
func PostUsers(c *gin.Context) {
	var newUser UserType

	// Call bindJSON to bind the received JSON to new Album
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	sqlStatement := `
	INSERT INTO user (age, email, first_name, last_name)
	VALUES (30, 'wonsang@freed.com', 'Wonsang', 'Chong'`
	VALUES (30, 'wonsang@freed.com', 'Wonsang', 'Chong'`

	// add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for an album
	// whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
*/
