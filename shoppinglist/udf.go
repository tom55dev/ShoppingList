package shoppinglist

var ShoppingItems = []ShoppingItem{}

var count = 1

// Queries

func GetShoppingItems() []ShoppingItem {
	return ShoppingItems
}

func GetShoppingItem(id int) ShoppingItem {
	for _, item := range ShoppingItems {
		if item.ID == id {
			return item
		}
	}
	return ShoppingItem{}
}

// Mutations

func AddShoppingItem(itemName string, description string, countItem int) *ShoppingItem {
	temp := &ShoppingItem{
		ID:          count,
		ItemName:    itemName,
		Description: description,
		Count:       countItem,
		Purchased:   false,
	}
	ShoppingItems = append(ShoppingItems, *temp)
	count++
	return temp
}

func UpdateShoppingItem(id int, itemName string, description string, countItem int, purchased bool) bool {
	for i, item := range ShoppingItems {
		if item.ID == id {
			ShoppingItems[i].ItemName = itemName
			ShoppingItems[i].Description = description
			ShoppingItems[i].Count = countItem
			ShoppingItems[i].Purchased = purchased
			return true
		}
	}
	return false
}

func DeleteShoppingItem(id int) bool {
	for i, item := range ShoppingItems {
		if item.ID == id {
			ShoppingItems = append(ShoppingItems[:i], ShoppingItems[i+1:]...)
			return true
		}
	}
	return false
}
