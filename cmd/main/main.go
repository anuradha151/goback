package main

import (
	"log"
	"net/http"

	"github.com/anuradha151/goback/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterAnnexRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
	
}

