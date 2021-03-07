package main

import (
	"log"

	"github.com/filipemunhoz/myprogram/helpers"
)

func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "Tipo 1"

	log.Println(myVar.TypeName)
}
