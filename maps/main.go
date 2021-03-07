package main

import "log"

type User struct {
	Name string
	Age  int
}

func main() {

	var myString string
	var myInt int

	myString = "Hi"
	myInt = 666

	mySecondString := "Another String"

	log.Println(myString, myInt, mySecondString)

	myMap := make(map[string]string)
	myMap["dog"] = "Dara"

	log.Println(myMap["dog"])

	me := User{
		Name: "John",
		Age:  73,
	}

	mapUsers := make(map[string]User)
	mapUsers["me"] = me
	log.Println(mapUsers["me"].Name)

}
