package graph

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShoppingListGraphHandler(t *testing.T) {
	query := `{
      "query": "{ shoppingItems { id itemName } }",
      "operation": ""
   }`
	req, err := http.NewRequest("POST", "/shoppinglist", bytes.NewBuffer([]byte(query)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShoppingListGraphHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
