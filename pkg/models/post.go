package models

import (
	"database/sql"
	"log"

	"github.com/anuradha151/goback/pkg/config"
)

var db *sql.DB

type post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	createPostTable(db)

}

func (p *post) createPost() *post {
	query := `INSERT INTO posts (title, body) 
	VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, p.Title, p.Body).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	p.ID = id
	return p
}

func findById(pk int) post {
	var id int
	var title string
	var body string

	query := `SELECT id, title, body FROM posts WHERE id = $1`
	err := db.QueryRow(query, pk).Scan(&id, &title, &body)

	if err != nil {
		log.Fatal(err)
	}

	return post{ID: id, Title: title, Body: body}

}

func findAll() []post {
	var posts []post
	query := `SELECT id, title, body FROM posts`
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var title string
	var body string

	for rows.Next() {
		err := rows.Scan(&id, &title, &body)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post{ID: id, Title: title, Body: body})
	}

	return posts

}

func createPostTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		body VARCHAR(1000) NOT NULL
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}


