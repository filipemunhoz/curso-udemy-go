package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberofTeeth int
}

func main() {

	dog := Dog{
		Name:  "Data",
		Breed: "Street Dog",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberofTeeth: 30,
	}

	PrintInfo(gorilla)

}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Uhuuul"
}

func (g Gorilla) NumberOfLegs() int {
	return 3
}

func PrintInfo(a Animal) {
	log.Println("This Animal says: ", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
