package main

import (
	"log"
	"net/http"

	"github.com/anuradha151/goback/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := mux.NewRouter()
	routes.RegisterAnnexRoutes(r)
	http.Handle("/", r)

	log.Println("Server started on: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))	
	
}

