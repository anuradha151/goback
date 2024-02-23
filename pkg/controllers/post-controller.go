package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anuradha151/goback/pkg/models"
)

var NewPost models.post 

func GetPosts(w http.ResponseWriter, r *http.Request) {
	newPosts := models.findAll()
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
