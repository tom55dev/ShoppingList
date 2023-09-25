package shoppinglist

import "testing"

func TestShoppingItem(t *testing.T) {
	item := &ShoppingItem{
		ID:          1,
		ItemName:    "Test Item",
		Description: "Test Description",
		Count:       2,
		Purchased:   false,
	}

	if item.ItemName != "Test Item" {
		t.Errorf("ItemName was incorrect, got: %s, want: %s", item.ItemName, "Test Item")
	}
}
