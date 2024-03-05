package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/anuradha151/goback/pkg/models"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	postID := vars["id"]
	id, err := strconv.Atoi(postID)
	if err != nil {
		log.Fatal(err)
	}
	post := models.FindById(id)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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

	_ = json.NewDecoder(r.Body).Decode(&NewPost)

	NewPost.UpdatePost()

	res, _ := json.Marshal(NewPost)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}
