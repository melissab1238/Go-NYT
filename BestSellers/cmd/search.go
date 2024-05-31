package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search books by keyword, author, or title",
	Long:  `Search books in the specified booklist by keyword, author, or title.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyword := cmd.Flag("keyword").Value.String()
		author := cmd.Flag("author").Value.String()
		title := cmd.Flag("title").Value.String()

		// Perform the search based on the provided flags
		var books []nytapi.Book
		if keyword != "" {
			books = searchBooksByKeyWord(keyword)
		} else if author != "" {
			books = searchBooksByAuthor(author)
		} else if title != "" {
			books = searchBooksByTitle(title)
		}

		// Display the results
		if len(books) > 0 {
			for _, book := range books {
				fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
			}
		} else {
			fmt.Println("No results found.")
		}
	},
}

func searchBooksByKeyWord(keyword string) []nytapi.Book {
	// Fetch all books
	fetchedBooks, err := FetchAllBooks()
	if err != nil {
		log.Fatalf("Failed to fetch books: %v", err)
	}

	// Implemneting fuzzy matching
	var matchedBooks []nytapi.Book

	dmp := diffmatchpatch.New()

	// Iterate over all books
	for _, book := range fetchedBooks {
		// Normalize the title and keyword to lower case
		normalizedTitle := strings.ToLower(book.Title)
		normalizedKeyword := strings.ToLower(keyword)

		// Perform fuzzy matching
		diffs := dmp.DiffMain(normalizedTitle, normalizedKeyword, false)

		// Calculate the similarity score
		similarityScore := calculateSimilarity(diffs)

		// fmt.Printf("Similarity score for '%s': %.2f\n", book.Title, similarityScore)

		// Define a threshold for considering a match
		threshold := 0.5 // Adjust based on your needs

		// If the similarity score exceeds the threshold, consider it a match
		if similarityScore >= threshold {
			matchedBooks = append(matchedBooks, book)
		}
	}
	return matchedBooks

}

// Helper function to calculate similarity score from diffs
func calculateSimilarity(diffs []diffmatchpatch.Diff) float64 {
	totalLength := 0
	matchLength := 0

	for _, diff := range diffs {
		totalLength += len(diff.Text)
		if diff.Type == diffmatchpatch.DiffInsert || diff.Type == diffmatchpatch.DiffEqual {
			matchLength += len(diff.Text)
		}
	}

	if totalLength == 0 {
		return 0
	}

	return float64(matchLength) / float64(totalLength)
}

func FetchAllBooks() ([]nytapi.Book, error) {
	listName := "hardcover-fiction" // hardcoding listName for now
	dateStr := "current"            // hardcoding current data for now

	// Encode the listName to ensure it's safe for inclusion in a URL
	encodedListName := strings.ReplaceAll(listName, " ", "%20")

	// Construct the URL to fetch all books for the given list
	url := fmt.Sprintf("https://api.nytimes.com/svc/books/v3/lists/%s/%s.json", dateStr, encodedListName)
	jsonData, err := nytapi.GetJsonFromUrl(url)
	if err != nil {
		return nil, err
	}

	var response nytapi.NYTResponse
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "OK" {
		return nil, fmt.Errorf(`response status is NOT "OK": %s`, response.Status)
	}

	return response.Results.Books, nil
}

func searchBooksByAuthor(author string) []nytapi.Book {
	// Implement the logic to search books by author
	// This could involve fetching all books and filtering based on the author
	return nil
}

func searchBooksByTitle(title string) []nytapi.Book {
	// Implement the logic to search books by title
	// This could involve fetching all books and filtering based on the title
	return nil
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Add persistent flags for global use
	rootCmd.PersistentFlags().StringP("keyword", "k", "", "Keyword to search for")

	// Local flags for this specific command
	searchCmd.Flags().StringP("author", "a", "", "Author to search for")
	searchCmd.Flags().StringP("title", "t", "", "Title to search for")
}
