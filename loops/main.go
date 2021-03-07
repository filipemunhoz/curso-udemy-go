package main

import "log"

func main() {

	for i := 0; i < 10; i++ {
		log.Println("result =", i)
	}

	bands := []string{"Iron Maiden", "Metallica", "Slayer", "Baroes da Pisadinha"}

	for _, x := range bands {

		log.Println(x)
	}

}
