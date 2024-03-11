package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrintJSON(jsonData []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonData, "", " ")
	if err != nil {
		log.Fatal("JSON marshaling failed: ", err)
	}
	fmt.Println(prettyJSON.String())
}
