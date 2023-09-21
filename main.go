package main

import (
	"net/http"
	"os"
	"shopping-list-app/graph"
)

func main() {
	http.HandleFunc("/shoppinglist", graph.ShoppingListGraphHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
