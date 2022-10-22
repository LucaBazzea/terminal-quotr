package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quote struct {
	Text [1]string `json:"Dwight_Schrute"`
}

func loadQuote(filename string) (Quote, error) {
	var quote Quote

	quoteFile, err := os.Open(filename)
	defer quoteFile.Close()

	if err != nil {
		return quote, err
	}

	jsonParser := json.NewDecoder(quoteFile)
	err = jsonParser.Decode(&quote)

	return quote, err
}

func main() {
	fmt.Println("Opened JSON file")

	quote, _ := loadQuote("schrute.json")

	fmt.Println(quote)
}
