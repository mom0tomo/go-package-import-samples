package main

import (
	"fmt"

	"github.com/mom0tomo/go-fizzbuzz"
)

func main() {

	for i := 0; i < 20; i++ {

		output := fizzbuzz.GetOutput(i)
		fmt.Println(output)
	}
}
