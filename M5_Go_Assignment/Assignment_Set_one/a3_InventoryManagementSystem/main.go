package main

import (
	"a3_InventoryManagementSystem/handlers"
	"a3_InventoryManagementSystem/model"
	"fmt"
)

func main() {
	Database := []model.InventoryStruct{
		{ID: 1, Name: "Chips", Price: 100, Stock: 8},
		{ID: 2, Name: "Phones", Price: 1000, Stock: 12},
		{ID: 3, Name: "Fruits", Price: 50, Stock: 18},
		{ID: 4, Name: "Sofa", Price: 8000, Stock: 12},
	}

	var option int

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Product Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Exit")

		fmt.Scan(&option)

		switch option {
		case 1:

			updatedDatabase, err := handlers.AddProduct(Database)
			if err != nil {
				fmt.Println("Error adding product:", err)
			} else {
				Database = updatedDatabase
				fmt.Println("Product added successfully.")
			}

		case 2:
			var id int
			fmt.Println("Enter the product ID to update stock:")
			fmt.Scan(&id)
			updatedDatabase, err := handlers.UpdateProduct(Database, id)
			if err != nil {
				fmt.Println("Error updating product stock:", err)
			} else {
				Database = updatedDatabase
				fmt.Println("Product stock updated successfully.")
			}

		case 3:
			var searchTerm string
			var searchOption int
			fmt.Println("Search by: 1. ID 2. Name")
			fmt.Scan(&searchOption)
			if searchOption == 1 {
				fmt.Println("Enter the product ID:")
				fmt.Scan(&searchTerm)
				product, err := handlers.SearchProduct(Database, searchTerm)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Product found:", *product)
				}
			} else if searchOption == 2 {
				fmt.Println("Enter the product Name:")
				fmt.Scan(&searchTerm)
				product, err := handlers.SearchProduct(Database, searchTerm)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Product found:", *product)
				}
			} else {
				fmt.Println("Invalid option.")
			}

		case 4:
			var sortOption string
			fmt.Println("Sort inventory by: 1. Price 2. Stock")
			fmt.Scan(&sortOption)
			if sortOption == "1" {
				handlers.DisplayInventory(Database, "price")
			} else if sortOption == "2" {
				handlers.DisplayInventory(Database, "stock")
			} else {
				fmt.Println("Invalid sort option.")
			}

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
