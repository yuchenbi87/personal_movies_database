package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite", "./movies.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := `
		CREATE TABLE IF NOT EXISTS movies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			year INTEGER,
			rank FLOAT
		);

		CREATE TABLE IF NOT EXISTS genres (
			movie_id INTEGER,
			genre TEXT
		);
	`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	populateMoviesTable(db, "./IMDB-movies.csv")
	populateGenresTable(db, "./IMDB-movies_genres.csv")

	// Example query
	query := `
		SELECT g.genre, AVG(m.rank) AS avg_rating
		FROM movies m
		JOIN genres g ON m.id = g.movie_id
		GROUP BY g.genre
		ORDER BY avg_rating DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Genre\tAverage Rating")
	for rows.Next() {
		var genre string
		var avgRating float64
		err = rows.Scan(&genre, &avgRating)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%.2f\n", genre, avgRating)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func populateMoviesTable(db *sql.DB, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.Default(int64(len(records) - 1))

	for _, record := range records[1:] {
		if len(record) < 4 {
			log.Printf("Skipping malformed record: %v", record)
			continue
		}
		_, err = db.Exec("INSERT OR REPLACE INTO movies (id, name, year, rank) VALUES (?, ?, ?, ?)", strings.TrimSpace(record[0]), strings.TrimSpace(record[1]), strings.TrimSpace(record[2]), strings.TrimSpace(record[3]))
		if err != nil {
			log.Printf("Error inserting record: %v, error: %v", record, err)
		}
		bar.Add(1)
	}
}

func populateGenresTable(db *sql.DB, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.Default(int64(len(records) - 1))

	for _, record := range records[1:] {
		if len(record) < 2 {
			log.Printf("Skipping malformed record: %v", record)
			continue
		}
		_, err = db.Exec("INSERT OR IGNORE INTO genres (movie_id, genre) VALUES (?, ?)", strings.TrimSpace(record[0]), strings.TrimSpace(record[1]))
		if err != nil {
			log.Printf("Error inserting record: %v, error: %v", record, err)
		}
		bar.Add(1)
	}
}
