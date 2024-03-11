package cmd

import (
	"fmt"
	"log"

	"github.com/melissab1238/GO-NYT/BestSellers/data"
	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Get all booklists from NYT",
	Long:  `All of the booklists offered by the NYT Bestsellers API`,
	Run: func(cmd *cobra.Command, args []string) {
		GetBookLists()
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)

}

func GetBookLists() []nytapi.BookList {
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
	return booklists
}

func PrintBooklistNames(bookLists []nytapi.BookList) {
	for _, bookList := range bookLists {
		fmt.Printf("%d %s\n", bookList.ID, bookList.ListName)
	}
}
