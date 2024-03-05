package routes

import (
	"github.com/anuradha151/goback/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAnnexRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")
}