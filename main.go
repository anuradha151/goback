package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type post struct {
	ID	  string `json:"id"` 
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []post{
	{ID: "1", Title: "Title 1", Body: "Body 1"},
	{ID: "2", Title: "Title 2", Body: "Body 2"},
	{ID: "3", Title: "Title 3", Body: "Body 3"},
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)
	router.POST("/posts", createPost)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
	router.Run("localhost:8080")
}

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	for _, p := range posts {
		if p.ID == id {
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
		if p.ID == id {
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
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}
