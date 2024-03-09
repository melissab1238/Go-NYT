package cli

import (
	"fmt"
	"log"

	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
)

type Command struct {
	Name        string
	Description string
	Execute     func()
}

var APIKEY string

func SetupCLI(apiKey string) {
	APIKEY = apiKey
	// TODO some more stuff
}

var Commands = map[string]Command{
	"list": {
		Name:        "list",
		Description: "List all best-selling lists",
		Execute:     displayBookLists,
	},
	"search": {
		Name:        "search",
		Description: "Search for books by title, author, or date",
		Execute:     searchBooks,
	},
	"hardcover": {
		Name:        "hardcover",
		Description: "Get hardcover books",
		Execute:     getHardcoverBooks,
	},
}

func displayBookLists() {
	// Fetch book lists
	bookLists, err := nytapi.FetchBookLists(APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	formatBookLists(bookLists)
}

func searchBooks() {
	fmt.Println("not yet implemented")
}

func getHardcoverBooks() {
	// Get Hardcover books
	books, err := nytapi.GetBooks(APIKEY, 2)
	if err != nil {
		log.Fatal(err)
	}
	// Display the books
	DisplayBooks(books)
}
