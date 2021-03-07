package main

import (
	"log"
	"sort"
)

func main() {

	var mySlice []string

	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Bob")
	mySlice = append(mySlice, "Ana")

	sort.Strings(mySlice)
	log.Println(mySlice)

	myInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(myInts)

	log.Println(myInts[0:2])
	log.Println(myInts[3:9])
}
