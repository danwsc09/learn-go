package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	HOST     string
	DATABASE string
	USER     string
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {
	config := DbConfig{HOST: "localhost", DATABASE: "testdb", USER: os.Getenv("DBUSER")}
	var connectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", config.HOST, config.USER, config.DATABASE)
	var err error

	// Get a database handle.
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("=================")
	fmt.Println("successfully connected", db)
	fmt.Println("=================")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
	fmt.Println("=================")

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	fmt.Println("=================")
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	// use $1, $2, ... instead of ? for PostgreSQL (vs MySQL)
	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID
func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	lastInsertId := 0
	row := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price)
	err := row.Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", alb)
	}

	return int64(lastInsertId), nil

	// MySQL way
	/*
		result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
		if err != nil {
			return 0, fmt.Errorf("addAlbum: %v", err)
		}

		// TODO: LastInsertId not supported by pg
		id, err := result.LastInsertId()

		if err != nil {
			return 0, fmt.Errorf("addAlbum: %v", err)
		}
		return id, nil
	*/
}
