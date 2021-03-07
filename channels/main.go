package main

import (
	"log"

	"github.com/filipemunhoz/myprogram/helpers"
)

const numPool = 10

// CalculateValue ...
func CalculateValue(intChan chan int) {

	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)

}
