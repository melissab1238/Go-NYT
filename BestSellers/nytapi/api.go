package nytapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func FetchBookLists() ([]BookList, error) {
	url := "https://api.nytimes.com/svc/books/v3/lists/names.json"
	jsonData, err := GetJsonFromUrl(url)
	if err != nil {
		return nil, err
	}

	// Assuming jsonData contains the JSON data as a byte slice
	var bookListFromAPI BookListFromAPI
	err = json.Unmarshal(jsonData, &bookListFromAPI)
	if err != nil {
		log.Fatal(err)
	}

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
func EncodeStringBase64(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func GetBestSellersByDate(date string, listName string) ([]Book, error) {
	// Encode the listName to ensure it's safe for inclusion in a URL
	encodedListName := strings.ReplaceAll(listName, " ", "%20")

	// bookListID is not being used
	url := fmt.Sprintf("https://api.nytimes.com/svc/books/v3/lists/%s/%s.json", date, encodedListName)
	jsonData, err := GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(nil, err)
	}

	var response NYTResponse
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		log.Fatal(nil, err)
	}

	if response.Status != "OK" {
		log.Fatal(nil, `Response status is NOT "OK`)
	}

	return response.Results.Books, nil
}
