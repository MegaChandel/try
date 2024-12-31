package handlers

import (
	"a3_InventoryManagementSystem/model"
	"fmt"
	"reflect"
	"sort"
)

func AddProduct(Database []model.InventoryStruct) ([]model.InventoryStruct, error) {

	var result model.InventoryStruct

	var NewName string
	var NewPrice float64
	var NewStock int

	result.ID = len(Database) + 2

	fmt.Println("Enter the name of the product")
	fmt.Scan(&result.Name)

	fmt.Println("Enter the price of the product")

	fmt.Scan(&result.Price)
	if reflect.TypeOf(result.Price) != reflect.TypeOf(float64(0)) {
		return nil, fmt.Errorf("error")
	}

	fmt.Println("Enter the amount of new stock product")
	fmt.Scan(&result.Stock)

	Database = append(Database, result)

	return Database, nil

}

func UpdateProduct(Database []model.InventoryStruct, id int) ([]model.InventoryStruct, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id cannot be zero")
	}
	var option int
	fmt.Println("Do you want \n 1.Add more stocks \n 2.Decrease the amount of stock")
	_, err := fmt.Scan(&option)
	if err != nil {
		return nil, fmt.Errorf("Cannot be a string ")

	}
	switch option {
	case 1:
		var amnt int
		fmt.Println("Enter the amount you want to add")
		_, err := fmt.Scan(&amnt)
		if amnt < 0 {
			return nil, fmt.Errorf("amnt cannot be less than 0")
		}
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		for i, product := range Database {
			if product.ID == id {
				Database[i].Stock = Database[i].Stock + amnt
			}
		}
	case 2:
		var amnt int
		fmt.Println("Enter the amount you want to deduct")
		_, err := fmt.Scan(&amnt)
		if amnt < 0 {
			return nil, fmt.Errorf("amnt cannot be less than 0")
		}
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		for i, product := range Database {
			if product.ID == id {
				Database[i].Stock = Database[i].Stock - amnt
			}
		}

	}
	return Database, nil

}

func SearchProduct(Database []model.InventoryStruct, searchTerm interface{}) (*model.InventoryStruct, error) {
	switch v := searchTerm.(type) {
	case int:
		for _, product := range Database {
			if product.ID == v {
				return &product, nil
			}
		}
		return nil, fmt.Errorf("product with ID %d not found", v)

	case string:
		for _, product := range Database {
			if product.Name == v {
				return &product, nil
			}
		}
		return nil, fmt.Errorf("product with name '%s' not found", v)

	default:
		return nil, fmt.Errorf("invalid search term type, use either int (ID) or string (Name)")
	}
}

func DisplayInventory(Database []model.InventoryStruct, sortBy string) {
    sortedInventory, err := SortInventory(Database, sortBy)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(sortedInventory) == 0 {
        fmt.Println("No products in inventory.")
        return
    }

    for _, product := range sortedInventory {
        fmt.Println(product.ID, product.Name, product.Stock, product.Price)
    }
}

func SortInventory(Database []model.InventoryStruct, sortBy string) ([]model.InventoryStruct, error) {
    sortedInventory := make([]model.InventoryStruct, len(Database))
    copy(sortedInventory, Database)

    switch sortBy {
    case "price":
        sort.Slice(sortedInventory, func(i, j int) bool {
            return sortedInventory[i].Price < sortedInventory[j].Price
        })
    case "stock":
        sort.Slice(sortedInventory, func(i, j int) bool {
            return sortedInventory[i].Stock < sortedInventory[j].Stock
        })
    default:
        return nil, fmt.Errorf("invalid sort option, use 'price' or 'stock'")
    }

    return sortedInventory, nil
}
