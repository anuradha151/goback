package main

import (
	"log"
	"net/http"

	"github.com/anuradha151/goback/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Starting the application")

	r := mux.NewRouter()
	routes.RegisterAnnexRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

