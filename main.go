package main

import (
	"fmt"
	"os"
)

// type author struct {
// 	author []author `json:"Dwight Schrute""`
// }

func main() {
	jsonFile, err := os.Open("schrute.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Opened JSON file")

	defer jsonFile.Close()

	fmt.Println(jsonFile)
}
