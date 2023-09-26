package main

import (
	"net/http"
	"os"
	"shopping-list-app/graph"
)

func main() {
	shoppingListHandler := func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Set the credentials

		// Call actual handler function
		graph.ShoppingListGraphHandler(w, r)
	}

	http.HandleFunc("/shoppinglist", shoppingListHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
