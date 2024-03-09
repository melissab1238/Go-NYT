package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/melissab1238/GO-NYT/BestSellers/cli"
	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get API key for NYT books
	apiKey := os.Getenv("API_KEY")

	// Initialize the CLI
	cli.RegisterCommands()

	// Fetch book lists
	bookLists, err := nytapi.FetchBookLists(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Display the book lists
	cli.DisplayBookLists(bookLists)

	// Get Hardcover books
	books, err := nytapi.GetBooks(apiKey, 2)
	if err != nil {
		log.Fatal(err)
	}
	// Display the books
	cli.DisplayBooks(books)
}
