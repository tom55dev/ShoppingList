package shoppinglist

import "testing"

func TestCRUDOperations(t *testing.T) {
	item := AddShoppingItem("Test Item", "Test Description", 2)
	if item.ItemName != "Test Item" {
		t.Errorf("ItemName was incorrect, got: %s, want: %s", item.ItemName, "Test Item")
	}

	updated := UpdateShoppingItem(item.ID, "Updated Item", "Updated Description", 3, true)
	if !updated {
		t.Errorf("UpdateShoppingItem was incorrect, got: %v, want: %v", updated, true)
	}

	deleted := DeleteShoppingItem(item.ID)
	if !deleted {
		t.Errorf("DeleteShoppingItem was incorrect, got: %v, want: %v", deleted, true)
	}
}

func TestUpdatePurchasedStatus(t *testing.T) {
	item := AddShoppingItem("Test Item", "Test Description", 2)
	updated := UpdatePurchasedStatus(item.ID, true)
	if !updated {
		t.Errorf("UpdatePurchasedStatus was incorrect, got: %v, want: %v", updated, true)
	}
}
