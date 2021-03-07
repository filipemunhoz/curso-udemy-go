package main

import "log"

func main() {
	var myString string
	myString = "Penca"

	log.Println("This is my String: ", myString)

	changeUsingPointer(&myString)

	log.Println("After change is my String: ", myString)
}

func changeUsingPointer(s *string) {

	log.Println("S is", s)
	log.Println("S is", *s)
	newValue := "Red"
	*s = newValue
}
