package nytapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// helper function
func getJsonFromUrl(url string, api_key string) ([]byte, error) {
	url = fmt.Sprintf("%s?api-key=%s", url, api_key)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FetchBookLists(apiKey string) ([]BookList, error) {
	url := "https://api.nytimes.com/svc/books/v3/lists/names.json"
	jsonData, err := getJsonFromUrl(url, apiKey)
	if err != nil {
	return nil, err
	}

	// Assuming jsonData contains the JSON data as a byte slice
	var bookListFromAPI BookListFromAPI
	err = json.Unmarshal(jsonData, &bookListFromAPI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of book lists:", bookListFromAPI.NumResults)

	// Add IDs to each booklist result
	var bookLists []BookList
	for _, bookList := range bookListFromAPI.Results {
		bookLists = append(bookLists,
			BookList{ID: len(bookLists) + 1,
				ListName:            bookList.ListName,
				DisplayName:         bookList.DisplayName,
				ListNameEncoded:     bookList.ListNameEncoded,
				OldestPublishedDate: bookList.OldestPublishedDate,
				NewestPublishedDate: bookList.NewestPublishedDate,
				Updated:             bookList.Updated})
	}
	return bookLists, nil
}

func GetBooks(apiKey string, bookListID int) ([]Book, error) {
	url := "https://api.nytimes.com/svc/books/v3/lists/current/hardcover-fiction.json"
	jsonData, err := getJsonFromUrl(url, apiKey)
	if err != nil {
		log.Fatal(nil, err)
	}

	var response NYTResponse
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		log.Fatal(nil, err)
	}
	// Access known status field
	status := response.Status
	fmt.Println("status:", status)
	// Access the number of books in the list
	numBooks := len(response.Results.Books)
	fmt.Println("Number of books in the list:", numBooks)

	return response.Results.Books, nil
}
