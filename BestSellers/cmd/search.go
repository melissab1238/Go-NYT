package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/melissab1238/GO-NYT/BestSellers/data"
	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var searchCmd = &cobra.Command{
	Use:   "listSearch",
	Short: "Search books in a booklist",
	Long:  `Search books in a booklist`,
	Run: func(cmd *cobra.Command, args []string) {
		listSearch()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// TODO use flags
	rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

func promptGetBooklist(pc promptContent) string {
	booklists := data.GetBookLists()
	var listNames []string
	for _, booklist := range booklists {
		listNames = append(listNames, booklist.ListName)
	}

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    listNames,
			AddLabel: "Other", // uhh i dont think i want an other option
		}

		index, result, err = prompt.Run()

		if index == -1 {
			listNames = append(listNames, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptGetBookTitle(pc promptContent, books []nytapi.Book) string {
	var bookNamesAndAuthors []string
	for _, book := range books {
		bookNamesAndAuthors = append(bookNamesAndAuthors, fmt.Sprintf("%-30s by %-30s", book.Title, book.Author))
	}

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    bookNamesAndAuthors,
			AddLabel: "Other", // TODO uhh i dont think i want an other option
		}

		index, result, err = prompt.Run()

		if index == -1 {
			bookNamesAndAuthors = append(bookNamesAndAuthors, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Split the result by " by" and take the first part with trailing whitespaces removed
	bookTitle := strings.Split(result, " by")[0]
	bookTitle = strings.TrimSpace(bookTitle)
	fmt.Printf("Input: %s\n", bookTitle)

	return bookTitle
}

func listSearch() {
	booklistPromptContent := promptContent{
		"Please provide a booklist.",
		"What booklist do you want to search from?",
	}
	listName := promptGetBooklist(booklistPromptContent)

	currentPromptContent := promptContent{
		`Please give a yes/no answer.`,
		"Do you want the most recent booklist data? (yes/no)",
	}
	current := promptGetInput(currentPromptContent)

	var books []nytapi.Book
	var err error
	if current == "yes" {
		books, err = nytapi.GetBestSellersByDate("current", listName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		yearPromptContent := promptContent{
			"Please provide a year (YYYY).",
			"What year would you like to search in? (YYYY)",
		}
		year := promptGetInput(yearPromptContent)

		monthPromptContent := promptContent{
			"Please provide a month (MM).",
			"What month would you like to search in? (MM)",
		}
		month := promptGetInput(monthPromptContent)

		date := fmt.Sprintf("%s-%s-01", year, month)
		books, err = nytapi.GetBestSellersByDate(date, listName)
		if err != nil {
			fmt.Println(err)
		}
	}

	bookPromptContent := promptContent{
		"Please pick a book.",
		"What book do you want to pick?",
	}
	bookTitle := promptGetBookTitle(bookPromptContent, books)

	book := FindBookByName(books, bookTitle)
	data.PrintBookInformation(book)
}

func FindBookByName(books []nytapi.Book, title string) *nytapi.Book {
	for _, book := range books {
		if book.Title == title {
			return &book
		}
	}
	return nil
}

// assistance from https://github.com/divrhino/studybuddy
