package main

import (
	"net/http"
	"os"
	"shopping-list-app/graph"
)

func main() {
	shoppingListHandler := func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

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
