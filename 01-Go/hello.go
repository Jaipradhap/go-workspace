package main

import (
	"fmt"

	"log"

	"one.com/greet"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello , Jai")
	fmt.Println(quote.Go())
	res, _ := greet.Greeting("Jai")
	fmt.Println(res)
	res, err := greet.Greeting("")
	// message := fmt.Sprintln(" %v - %v", res, )
	fmt.Println(err)
	fmt.Println(res)

	// Logs

	log.SetPrefix("Main : ")
	log.SetFlags(0)

	res, err = greet.Greeting("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
