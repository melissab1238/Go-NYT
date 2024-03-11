package data

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
)

var booklists []nytapi.BookList

func GetBookLists() []nytapi.BookList {
	if booklists != nil {
		return booklists
	}
	// read booklists from JSON file
	filePath := "./data/booklists.json"
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File does not exist: %s", filePath)
			return nil
		} else {
			log.Printf("Error checking file: %v", err)
			return nil
		}
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &booklists)
	return booklists
}

func SetBooklists(bl []nytapi.BookList) {
	booklists = bl
	saveToJSON(booklists)
}

func saveToJSON(bl []nytapi.BookList) {
	jsonData, err := json.MarshalIndent(bl, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling booklists to JSON: %v", err)
	}

	// Define the file path where the JSON will be saved
	filePath := "./data/booklists.json"

	// Write the JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}

	log.Printf("Booklists saved to %s", filePath)

}
