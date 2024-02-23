package routes

import (
	"github.com/gorilla/mux"
	"github.com/anuradha151/goback/pkg/controllers"
)

var RegisterAnnexRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")
}