package cli

import (
	"fmt"
	"log"

	"github.com/melissab1238/GO-NYT/BestSellers/data"
	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
)

type Command struct {
	Name        string
	Description string
	Execute     func()
}

var Commands = map[string]Command{
	"list": {
		Name:        "list",
		Description: "List all best-selling lists",
		Execute:     DisplayBookLists,
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

// Fetch book lists from api and print booklists
func DisplayBookLists() {
	var booklists []nytapi.BookList
	booklists = data.GetBookLists()
	if booklists == nil {

		var err error
		booklists, err = nytapi.FetchBookLists()
		if err != nil {
			log.Fatal(err)
		}
		data.SetBooklists(booklists)
	}
	// todo - add booklists to cache
	formatBookLists(booklists)
}

func searchBooks() {
	fmt.Println("not yet implemented")
}

func getHardcoverBooks() {
	// Get Hardcover books
	books, err := nytapi.GetBooks(2) // 2 is the index of hardcover books
	if err != nil {
		log.Fatal(err)
	}
	// Display the books
	DisplayBooks(books)
}
