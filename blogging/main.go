package main

import (
	"log"
	"net/http"

	"github.com/aanchalverma/Machine-Coding/blogging/routes"
)

func main() {
	// API to register the user
	http.HandleFunc("/register", routes.Register)
	// API for user login
	http.HandleFunc("/login", routes.Login)
	// API for user operations - Create/Update/Delete/Get
	http.HandleFunc("/posts", routes.PostsHandler)
	http.HandleFunc("/posts/", routes.PostHandler)

	// Start the host server on 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
