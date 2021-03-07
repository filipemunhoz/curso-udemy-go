package main

import (
	"log"
	"time"
)

var s = "seven"

// User ...
type User struct {
	FirstName string
	LastName  string
	Age       int
	Birthday  time.Time
}

func main() {

	user := User{
		FirstName: "Joseph",
		LastName:  "Silva",
		Age:       36,
		Birthday:  time.Date(1980, 01, 02, 20, 34, 58, 651387237, time.UTC),
	}

	log.Println(user.FirstName)
	log.Println(user.Birthday)
}
