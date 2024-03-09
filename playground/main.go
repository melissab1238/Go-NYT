package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	firstName, lastName, age := "Bob", "Sal", 100
	const publisher = "Penguin"
	fmt.Println(publisher)
	_, _, _ = firstName, lastName, age // compiler is now happy

	address := "123 Main St\nMinneapolis\nMinnesota, USA\n12345"
	_ = address
	fmt.Println(address)
	fmt.Printf("number of characters in address: %d\n", len(address))

	start := time.Now()
	fmt.Println(start)
	fmt.Printf("type of start: %s\n", reflect.TypeOf((start)))

	// CLI practice
	type Person struct {
		FirstName string
		LastName  string
		Age       int
		FavColor  string
	}
	var person Person
	// whats your name?
	fmt.Println("What's your first name?")
	fmt.Scanln(&person.FirstName)
	fmt.Println("What's your last name?")
	fmt.Scanln(&person.LastName)
	// whats your age?
	fmt.Println("What's your age?")
	fmt.Scanln(&person.Age)
	// whats your favorite color?
	fmt.Println("What's your favorite color?")
	fmt.Scanln(&person.FavColor)
	// save to struct, then save in json file

	// Marshall the struct into JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// Write the JSON data to a file
	err = os.WriteFile("person.json", jsonData, 0644) // 0644 means read and write for owner and read for everyone else
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
