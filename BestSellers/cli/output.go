package cli

import (
	"fmt"

	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
)

// DisplayBookLists displays the book lists to the console.
func DisplayBookLists(bookLists []nytapi.BookList) {
	for _, bookList := range bookLists {
		fmt.Printf("%d %s\n", bookList.ID, bookList.ListName)
	}
}

// DisplayBooks displays the books to the console.
func DisplayBooks(books []nytapi.Book) {
	for _, book := range books {
		fmt.Printf("%s\n", book.Title)
	}
}
