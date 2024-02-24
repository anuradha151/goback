package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anuradha151/goback/pkg/models"
)

var NewPost models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	newPosts := models.FindAll()
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	_ = json.NewDecoder(r.Body).Decode(&NewPost)
	NewPost.CreatePost()
	res, _ := json.Marshal(NewPost)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
}
