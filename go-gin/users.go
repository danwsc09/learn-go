package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user represents data about a record album.
type UserType struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users []UserType = []UserType{
	{Age: 10, Email: "tempaion1@hotmail.com", FirstName: "A1", LastName: "B1"},
	{Age: 20, Email: "tempaion2@hotmail.com", FirstName: "A2", LastName: "B2"},
	{Age: 30, Email: "tempaion3@hotmail.com", FirstName: "A3", LastName: "B3"},
	{Age: 40, Email: "tempaion4@hotmail.com", FirstName: "A4", LastName: "B4"},
	{Age: 50, Email: "tempaion5@hotmail.com", FirstName: "A5", LastName: "B5"},
	{Age: 60, Email: "tempaion6@hotmail.com", FirstName: "A6", LastName: "B6"},
	{Age: 70, Email: "tempaion7@hotmail.com", FirstName: "A7", LastName: "B7"},
	{Age: 80, Email: "tempaion8@hotmail.com", FirstName: "A8", LastName: "B8"},
	{Age: 90, Email: "tempaion9@hotmail.com", FirstName: "A9", LastName: "B9"},
	{Age: 110, Email: "tempaion11@hotmail.com", FirstName: "A11", LastName: "B11"},
	{Age: 120, Email: "tempaion12@hotmail.com", FirstName: "A12", LastName: "B12"},
	{Age: 130, Email: "tempaion13@hotmail.com", FirstName: "A13", LastName: "B13"},
	{Age: 140, Email: "tempaion14@hotmail.com", FirstName: "A14", LastName: "B14"},
	{Age: 150, Email: "tempaion15@hotmail.com", FirstName: "A15", LastName: "B15"},
	{Age: 160, Email: "tempaion16@hotmail.com", FirstName: "A16", LastName: "B16"},
	{Age: 170, Email: "tempaion17@hotmail.com", FirstName: "A17", LastName: "B17"},
	{Age: 180, Email: "tempaion18@hotmail.com", FirstName: "A18", LastName: "B18"},
	{Age: 190, Email: "tempaion19@hotmail.com", FirstName: "A19", LastName: "B19"},
}

// GetUsers responds with the list of all albums as JSON.
func GetUsers(c *gin.Context) {
	queryParam := c.Request.URL.Query()
	pageSize, pageIndex := queryParam.Get("page_size"), queryParam.Get("page_index")
	pageSizeI, _ := strconv.Atoi(pageSize)
	pageIndexI, _ := strconv.Atoi(pageIndex)
	start, end := pageSizeI*pageIndexI, pageSizeI*(pageIndexI+1)-1
	res := users[start : end+1]
	c.IndentedJSON(http.StatusOK, res)
	/*
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
	*/
}

func PostUsers(c *gin.Context) {
	var newUser UserType

	// Call bindJSON to bind the received JSON to new Album
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
	/*
		sqlStatement := `
		INSERT INTO user (age, email, first_name, last_name)
		VALUES (30, 'wonsang@freed.com', 'Wonsang', 'Chong'`
		// VALUES (30, 'wonsang@freed.com', 'Wonsang', 'Chong'`

		// add the new album to the slice
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	*/
}

func GetUsersFilterByAge(c *gin.Context) {
	ageParam := c.Param("minAge")
	minAge, err := strconv.Atoi(ageParam)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no users found"})
	}

	// loop over the list of albums, looking for an album
	// whose ID value matches the parameter
	var res []UserType
	for _, a := range users {
		if a.Age > minAge {
			res = append(res, a)
		}
	}
	c.IndentedJSON(http.StatusOK, res)
}

/*
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
