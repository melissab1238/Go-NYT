package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type BookListFromAPI struct {
	Status     string `json:"status"`
	Copyright  string `json:"copyright"`
	NumResults int    `json:"num_results"`
	Results    []struct {
		ListName            string `json:"list_name"`
		DisplayName         string `json:"display_name"`
		ListNameEncoded     string `json:"list_name_encoded"`
		OldestPublishedDate string `json:"oldest_published_date"`
		NewestPublishedDate string `json:"newest_published_date"`
		Updated             string `json:"updated"`
	} `json:"results"`
}

type BookList struct {
	ID                  int
	ListName            string
	DisplayName         string
	ListNameEncoded     string
	OldestPublishedDate string
	NewestPublishedDate string
	Updated             string
}

type Book struct {
	Rank               int       `json:"rank"`
	RankLastWeek       int       `json:"rank_last_week"`
	WeeksOnList        int       `json:"weeks_on_list"`
	Asterisk           int       `json:"asterisk"`
	Dagger             int       `json:"dagger"`
	PrimaryISBN10      string    `json:"primary_isbn10"`
	PrimaryISBN13      string    `json:"primary_isbn13"`
	Publisher          string    `json:"publisher"`
	Description        string    `json:"description"`
	Price              string    `json:"price"`
	Title              string    `json:"title"`
	Author             string    `json:"author"`
	Contributor        string    `json:"contributor"`
	ContributorNote    string    `json:"contributor_note"`
	BookImage          string    `json:"book_image"`
	BookImageWidth     int       `json:"book_image_width"`
	BookImageHeight    int       `json:"book_image_height"`
	AmazonProductURL   string    `json:"amazon_product_url"`
	AgeGroup           string    `json:"age_group"`
	BookReviewLink     string    `json:"book_review_link"`
	FirstChapterLink   string    `json:"first_chapter_link"`
	SundayReviewLink   string    `json:"sunday_review_link"`
	ArticleChapterLink string    `json:"article_chapter_link"`
	ISBNs              []ISBN    `json:"isbns"`
	BuyLinks           []BuyLink `json:"buy_links"`
	BookURI            string    `json:"book_uri"`
}

type ISBN struct {
	ISBN10 string `json:"isbn10"`
	ISBN13 string `json:"isbn13"`
}

type BuyLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NYTResponse struct {
	Status  string `json:"status"`
	Results struct {
		ListName                 string        `json:"list_name"`
		ListNameEncoded          string        `json:"list_name_encoded"`
		BestsellersDate          string        `json:"bestsellers_date"`
		PublishedDate            string        `json:"published_date"`
		PublishedDateDescription string        `json:"published_date_description"`
		NextPublishedDate        string        `json:"next_published_date"`
		PreviousPublishedDate    string        `json:"previous_published_date"`
		DisplayName              string        `json:"display_name"`
		NormalListEndsAt         int           `json:"normal_list_ends_at"`
		Updated                  string        `json:"updated"`
		Books                    []Book        `json:"books"`
		Corrections              []interface{} `json:"corrections"`
	} `json:"results"`
}

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

var listNames []string

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get API key for NYT books
	api_key := os.Getenv("API_KEY")

	// Get all list names
	url := "https://api.nytimes.com/svc/books/v3/lists/names.json"
	jsonData, err := getJsonFromUrl(url, api_key)
	if err != nil {
		log.Fatal(err)
	}
	// save to file
	err = os.WriteFile("list_name.json", jsonData, 0644) // 0644 means read and write for owner and read for everyone else
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
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

	// Now bookListsWithID contains all the book lists with their IDs
	// Output the bookListsWithID
	for _, bookList := range bookLists {
		fmt.Printf("%d %s\n", bookList.ID, bookList.ListName)
	}

	// Next section
	// Get hardcover fiction lists
	url = "https://api.nytimes.com/svc/books/v3/lists/current/hardcover-fiction.json"
	jsonData, err = getJsonFromUrl(url, api_key)
	if err != nil {
		log.Fatal(err)
	}

	var response NYTResponse
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		log.Fatal(err)
	}
	// Access known status field
	status := response.Status
	fmt.Println("status:", status)
	// Access the number of books in the list
	numBooks := len(response.Results.Books)
	fmt.Println("Number of books in the list:", numBooks)

}
