package shoppinglist

// ShoppingItem is a struct that represents an item in the shopping list
type ShoppingItem struct {
	ID          int    `json:"id"`
	ItemName    string `json:"itemName"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	Purchased   bool   `json:"purchased,omitempty"`
}
