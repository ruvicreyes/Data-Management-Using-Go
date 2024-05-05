package main

import (	
	"log"
	"encoding/json"
	"os"
)

func createNewDB(){
	// Create an empty slice ([]Person{}) to represent the data
	data := []interface{}{}

	// Marshal the empty slice to JSON format
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Write the JSON data to a file named "data.json"
	err = os.WriteFile("database.json", jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func writeData (newData []Product){

	//indent the Data
	jsonData, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		log.Print(err)
	}
	//Write data
	err = os.WriteFile("database.json", jsonData, 0644)
	if err != nil {
		log.Print(err)
	}
}

// Define a function to load data from a JSON file
func loadData(filename string) ([]Product, error) {
	var data []Product

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Function to check if a person already exists in the slice
func personExists(product []Product, newProduct Product) bool {
	for _, p := range product {
		if p.ID == newProduct.ID && p.Name == newProduct.Name && p.Price == newProduct.Price {
			return true
		}
	}
	return false
}


/*this will determine the path of write or read, 
can be use this feature in the future  */
// func databasePath(segments ...string) string {
// 	//Create a json database
// 	home, err := os.UserHomeDir()
// 	if err != nil {
// 		panic(err)
// 	}
// 	segments = append([]string{home, "go-concurr-pattern"}, segments...)
// 	return filepath.Join(segments...)
// }

