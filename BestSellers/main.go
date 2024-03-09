package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/melissab1238/GO-NYT/BestSellers/cli"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get API key for NYT books
	apiKey := os.Getenv("API_KEY")

	// Initialize the CLI
	cli.SetupCLI(apiKey)

	// Start the CLI
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Display book lists")
		fmt.Println("2. Search for books")
		fmt.Println("3. Get Hardcover book list")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\n":
			fmt.Println("Displaying book lists...")
			commandName := "list" // os.Args[1]
			command, exists := cli.Commands["list"]
			if !exists {
				fmt.Printf("Unknown command: %s\n", commandName)
				return
			}
			fmt.Println(command.Description)
			command.Execute()
		case "2\n":
			fmt.Println("Searching for books...")
			// Call your function to search for books here
		case "3\n":
			fmt.Println("Getting hard cover list...")
			commandName := "hardcover" // os.Args[1]
			command, exists := cli.Commands["hardcover"]
			if !exists {
				fmt.Printf("Unknown command: %s\n", commandName)
				return
			}
			fmt.Println(command.Description)
			command.Execute()
		case "4\n":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
