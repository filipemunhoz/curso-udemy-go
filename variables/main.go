package main

import (
	"log"
)

func main() {

	var whatToSay string = saySomething("Balanca brasil")
	var anotherElse string = saySomething("Bla bla bla")
	var firstReturn string
	var secondReturn string
	var i int = 8

	log.Println(whatToSay)
	log.Println(anotherElse)
	log.Println(i)

	firstReturn, secondReturn = withTwoValuesReturn("Aloha")
	log.Println(firstReturn, secondReturn)

	firstReturn, _ = withTwoValuesReturn("Welcome")
	log.Println(firstReturn)
}

func saySomething(s string) string {
	return s
}

func withTwoValuesReturn(s string) (string, string) {
	return s, "Hello"
}
