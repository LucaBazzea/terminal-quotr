package main

import (
	"fmt"
	"os"
)


func main() {
  dir, error := os.ReadDir("quotes")

  if error != nil {
    fmt.Println("Error: ", error)
    return
  }
  
  for _, entry := range dir {
    fmt.Println(entry.Name())
  }
}
