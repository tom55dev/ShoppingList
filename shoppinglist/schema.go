package shoppinglist

import "github.com/graphql-go/graphql"

// Schema for ShoppingItem
var shoppingItemSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ShoppingItem",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"itemName": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
			"purchased": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

// Queries
var shoppingItemQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ShoppingItemQueries",
		Fields: graphql.Fields{
			"shoppingItems": &graphql.Field{
				Type: graphql.NewList(shoppingItemSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetShoppingItems(), nil
				},
			},
			"shoppingItem": &graphql.Field{
				Type: shoppingItemSchema,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return GetShoppingItem(id), nil
					}
					return nil, nil
				},
			},
		},
	},
)

// Mutations
var shoppingItemMutations = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ShoppingItemMutations",
		Fields: graphql.Fields{
			"addShoppingItem": &graphql.Field{
				Type: shoppingItemSchema,
				Args: graphql.FieldConfigArgument{
					"itemName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"count": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					itemName, _ := p.Args["itemName"].(string)
					description, _ := p.Args["description"].(string)
					count, _ := p.Args["count"].(int)
					return AddShoppingItem(itemName, description, count), nil
				},
			},
			"updateShoppingItem": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"count": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"purchased": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					itemName, _ := p.Args["itemName"].(string)
					description, _ := p.Args["description"].(string)
					count, _ := p.Args["count"].(int)
					purchased, _ := p.Args["purchased"].(bool)
					return UpdateShoppingItem(id, itemName, description, count, purchased), nil
				},
			},
			"deleteShoppingItem": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return DeleteShoppingItem(id), nil
					}
					return false, nil
				},
			},
		},
	},
)

// Root Schema
var ShoppingListRootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    shoppingItemQueries,
		Mutation: shoppingItemMutations,
	},
)
