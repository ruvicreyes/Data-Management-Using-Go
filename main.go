package main

import (
	"log"
)

type Storer interface {
	AddProduct(product *Product)
	RemoveProduct(id int)
	//UpdateProduct(id int, product *Product)
	//GetProduct(id int) *Product
	GetAllProducts()
}

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (p Product) AddProduct(product *Product) {
	// Load existing data from the JSON file

	existingData, err := loadData("database.json")
	if err != nil {
		log.Print(err, "\n We Create New one called database.json")
		createNewDB()
	}
	if personExists(existingData, *product) {
		//check if data is already Existing
		log.Println("Person already exists, not adding.")

	} else {
		//Append the Existing Data to the new Data
		newData := append(existingData, *product)
		log.Println(newData)
		writeData(newData)
	}
}

func (p Product) RemoveProduct(id int) {

	existingData, err := loadData("database.json")
	if err != nil {
		panic(err)
	}

	var prod []Product
	// Iterate over the items and remove the item with the specified ID
	found := false
	for _, item := range existingData {
		if item.ID == id {
			found = true
			log.Println("ID NO:", id, "is completely removed")
		} else {
			prod = append(prod, item)
		}
	}
	// wew
	// If the item was not found, return an error
	if !found {
		log.Println("No ID Found")
	}

	//Update Database
	writeData(prod)

}

// func (p Product) updatedProduct(product *Product) {

// }

// func (p Product) GetProduct(id int) *Product {
// 	var list []string
//     for _, user := range *Product {
//         list = append(list, user.UserName)
//     }
//     return list

// }

func (p Product) GetAllProducts() {
	data, err := loadData("database.json")
	if err != nil {
		panic(err)
	}

	for _, p := range data {
		log.Println(p)
	}
}

func main() {
	//create Dir of the project
	// err := os.MkdirAll(databasePath(), 0700)
	// if err != nil {
	// 	panic(err)
	// }

	// Initialize the inventory management system
	inventory := Product{}

	// Add some products
	// inventory.AddProduct(&Product{ID: 1, Name: "Laptop", Price: 1000, Quantity: 10})
	// inventory.AddProduct(&Product{ID: 2, Name: "Smartphone", Price: 500, Quantity: 20})
	// inventory.AddProduct(&Product{ID: 3, Name: "tablet", Price: 1500, Quantity: 20})

	// List all products
	log.Println("All Products:")
	inventory.GetAllProducts()

	// // Remove a product
	inventory.RemoveProduct(1)

	// // Update a product
	// updatedProduct := &Product{ID: 2, Name: "Updated Smartphone", Price: 550, Quantity: 25}
	// inventory.UpdateProduct(2, updatedProduct)

	// // Get a product by ID
	// product := inventory.GetProduct(2)
	// fmt.Println("Updated Product:", product)
}

//Next is to update product

/* get list of product if == id
https://stackoverflow.com/questions/34172001/how-to-retrieve-array-of-elements-from-array-of-structure-in-golang */

// possible if make interface of writing data in Filemanipulation

/*
Benefits of this Project:

It provides hands-on experience with defining and implementing interfaces in Go.
It involves working with pointers to manipulate and update data structures.
It helps in understanding how to abstract functionality using interfaces, making the code more modular and extensible.
It allows practicing basic CRUD operations and managing data using in-memory storage.
It provides a practical context to apply concepts like composition, encapsulation, and abstraction in Go programming.
Feel free to expand upon this project by adding more features like persistence using databases, validation, error handling, or implementing a web-based interface using frameworks like Gin or Echo.
*/
