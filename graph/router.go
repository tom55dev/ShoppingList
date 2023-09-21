package graph

import (
	"encoding/json"
	"net/http"
	"shopping-list-app/shoppinglist" // replace with your project import path

	"github.com/graphql-go/graphql"
)

type RequestParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func ShoppingListGraphHandler(w http.ResponseWriter, r *http.Request) {
	var reqObj RequestParams

	if err := json.NewDecoder(r.Body).Decode(&reqObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// You can add more logic here like checking authentication, authorization, etc.

	result := graphql.Do(graphql.Params{
		Context:        r.Context(),
		Schema:         shoppinglist.ShoppingListRootSchema,
		RequestString:  reqObj.Query,
		VariableValues: reqObj.Variables,
		OperationName:  reqObj.Operation,
	})

	if len(result.Errors) > 0 {
		http.Error(w, result.Errors[0].Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
