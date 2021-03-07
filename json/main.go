package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person ...
type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {

	myJSON := `
[
  {
    "name": "Filipe",
    "last_name": "Maiden"
  },
  {
    "name": "John",
    "last_name": "Silva"
  }
]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJSON), &unmarshalled)

	if err != nil {
		log.Println("Error parsing json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person
	m1.Name = "Jack"
	m1.LastName = "Tequila"

	var m2 Person
	m2.Name = "Baranga"
	m2.LastName = "Metal"

	mySlice = append(mySlice, m2)

	newJSON, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error creating json", err)
	}

	fmt.Println(string(newJSON))

}
