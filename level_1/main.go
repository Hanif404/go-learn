package main

import (
	"errors"
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Hello, world!")

	const name, age = "Hanif", 30
	s := fmt.Sprintln(name, "is", age, "years old.")
	fmt.Print(s)

	e := fmt.Errorf("%s", "data not found").Error()
	fmt.Print(e)

	err := errors.New("failed get data")
	fmt.Print(err)
}
