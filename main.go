package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []post{
	{ID: 1, Title: "Title 1", Body: "Body 1"},
	{ID: 2, Title: "Title 2", Body: "Body 2"},
	{ID: 3, Title: "Title 3", Body: "Body 3"},
}

func main() {

	connStr := "postgres://root:root@localhost/annex?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createPostTable(db)
	pk := insertPost(db, post{Title: "Double room", Body: "New furniture, new tile, new bathwares"})
	log.Println("Primary key:", pk)

	p := findById(db, pk)
	log.Println("Post:", p)

	posts := findAllPosts(db)
	log.Println("Posts:", posts)

	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/post/:id", getPost)
	router.POST("/post", createPost)
	router.PUT("/post/:id", updatePost)
	router.DELETE("/post/:id", deletePost)
	router.Run("localhost:8080")
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	for _, p := range posts {
		if p.ID == strToInt(id) {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

func createPost(c *gin.Context) {
	var newPost post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedPost post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	for i, p := range posts {
		if p.ID == strToInt(id) {
			posts[i] = updatedPost
			c.JSON(http.StatusOK, updatedPost)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	for i, p := range posts {
		if p.ID == strToInt(id) {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
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

func insertPost(db *sql.DB, p post) int {
	query := `INSERT INTO posts (title, body) 
	VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, p.Title, p.Body).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	return id
}

func findById(db *sql.DB, pk int) post {
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

func findAllPosts(db *sql.DB) []post {
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

