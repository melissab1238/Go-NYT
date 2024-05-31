package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/melissab1238/GO-NYT/BestSellers/data"
	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Get all booklists from NYT",
	Long:  `All of the booklists offered by the NYT Bestsellers API`,
	Run: func(cmd *cobra.Command, args []string) {
		booklists := GetBookLists()
		PrintBooklistNames(booklists)
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

	// Create a new tablewriter
	table := tablewriter.NewWriter(os.Stdout)

	// Set table headers
	table.SetHeader([]string{"ID", "List Name", "Oldest Pub. Date", "Newest Pub. Data"})

	// Convert bookLists to the required data format
	data := [][]string{}
	for _, bookList := range bookLists {
		row := []string{strconv.Itoa(bookList.ID), bookList.ListName, bookList.OldestPublishedDate, bookList.NewestPublishedDate}
		data = append(data, row)
	}

	// Append data rows
	for _, v := range data {
		table.Append(v)
	}

	// Render the table
	table.Render()
}
