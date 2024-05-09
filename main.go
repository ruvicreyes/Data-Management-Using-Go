package main

import (
	"log"
)

type Storer interface {
	AddProduct(product *Product)
	RemoveProduct(id int)
	UpdateProduct(id int, product *Product)
	GetProduct(id int) *Product
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
		log.Println(err)
	}
	isPositive, ID := personExists(existingData, *product)
	if isPositive {
		//check if data is already Existing
		log.Println("The Product Number:", ID, "already exist")

	} else {
		//Append the Existing Data to the new Data
		newData := append(existingData, *product)
		writeData(newData)
	}
}

func (p Product) RemoveProduct(id int) {

	existingData, err := loadData("database.json")
	if err != nil {
		log.Println(err)
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

func (p Product) UpdateProduct(id int, product *Product) {
	var prod []Product
	found := false
	//load existing data
	existData, err := loadData("database.json")
	if err != nil {
		log.Println(err)
	}
	//Check if the Data is already updated
	isPositive, _ := alreadyUpdated(existData, *product)

	if isPositive {
		log.Println("The Item is already updated")
	} else {

		// Iterate over the items and remove the item with the specified ID
		for _, item := range existData {
			if item.ID == id {
				log.Println("ID NO:", id, "is completely updated")
				prod = append(prod, *product)
			} else {
				prod = append(prod, item)
				found = true
			}
		}
		// If the item was not found, return an error
		if !found {
			log.Println("No ID Found")
		}
		//Update Database
		writeData(prod)
	}

}

func (p Product) GetProduct(id int) (*Product, error) {
	isPositive := false
	//load existing data
	existData, err := loadData("database.json")
	if err != nil {
		log.Println(err)
	}
	// Iterate over the items and remove the item with the specified ID
	for _, item := range existData {
		if item.ID == id {
			isPositive = true
			return &item, err
		}
	}
	// If the item was not found, return an errorv
	if !isPositive {
		log.Println("No ID Found")

	}
	return nil, err
}

func (p Product) GetAllProducts() {
	data, err := loadData("database.json")
	if err != nil {
		log.Println(err)
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
	inventory.AddProduct(&Product{ID: 1, Name: "Laptop", Price: 1000, Quantity: 10})
	inventory.AddProduct(&Product{ID: 2, Name: "Smartphone", Price: 500, Quantity: 20})
	inventory.AddProduct(&Product{ID: 3, Name: "tablet", Price: 1500, Quantity: 20})

	// List all products
	log.Println("\nAll Products:")
	inventory.GetAllProducts()

	// // Remove a product
	inventory.RemoveProduct(1)

	// // Update a product
	updatedProduct := &Product{ID: 2, Name: "Updated Smartphone", Price: 550, Quantity: 25}
	inventory.UpdateProduct(2, updatedProduct)

	// // Get a product by ID
	product, err := inventory.GetProduct(2a)
	if product == nil {
		panic(err)
	} else {
		log.Println("Updated Product:", product)
	}
}

// possible additional feature is GORM before learning PostreSQL

/* possible if make interface of writing data in Filemanipulation
Writingfile
*/

/*
Benefits of this Project:

It provides hands-on experience with defining and implementing interfaces in Go.
It involves working with pointers to manipulate and update data structures.
It helps in understanding how to abstract functionality using interfaces, making the code more modular and extensible.
It allows practicing basic CRUD operations and managing data using in-memory storage.
It provides a practical context to apply concepts like composition, encapsulation, and abstraction in Go programming.
Feel free to expand upon this project by adding more features like persistence using databases, validation, error handling, or implementing a web-based interface using frameworks like Gin or Echo.
*/
