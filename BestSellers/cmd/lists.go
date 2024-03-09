/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/melissab1238/GO-NYT/BestSellers/cli"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Get all booklists from NYT",
	Long:  `All of the booklists offered by the NYT Bestsellers API`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.DisplayBookLists()
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)

}
