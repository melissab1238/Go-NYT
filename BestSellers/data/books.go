package data

import (
	"fmt"
	"os"

	"github.com/melissab1238/GO-NYT/BestSellers/nytapi"
	"github.com/olekukonko/tablewriter"
)

func PrintBookInformation(book *nytapi.Book) {
	// Create a new tablewriter
	table := tablewriter.NewWriter(os.Stdout)

	// Set table options
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Align all columns to the left
	table.SetRowLine(true)                     // Enable row lines

	// Add the column headers and book data as separate rows within two columns
	table.Append([]string{"Title", book.Title})
	table.Append([]string{"Author", book.Author})
	table.Append([]string{"Description", book.Description})
	table.Append([]string{"Price", book.Price})
	table.Append([]string{"Contributor", book.Contributor})
	table.Append([]string{"Publisher", book.Publisher})
	table.Append([]string{"Age Group", book.AgeGroup})
	table.Append([]string{"Weeks on List", fmt.Sprintf("%d", book.WeeksOnList)})
	table.Append([]string{"Rank", fmt.Sprintf("%d", book.Rank)})

	// Render the table
	table.Render()

}
