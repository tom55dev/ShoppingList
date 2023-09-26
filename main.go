package main

import (
	"log"
	"net/http"
	"os"
	"shopping-list-app/graph"

	"github.com/rs/cors"
)

func main() {
	shoppingListHandler := func(w http.ResponseWriter, r *http.Request) {

		// Log the request
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		// Call actual handler function
		graph.ShoppingListGraphHandler(w, r)
	}

	mux := http.NewServeMux()

	// Replace this with your actual API setup.
	mux.HandleFunc("/shoppinglist", shoppingListHandler)

	// Use default options
	handler := cors.Default().Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)

	http.ListenAndServe(":"+port, handler)
}
