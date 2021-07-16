package main

import (
	"fmt"

	"github.com/cameronperera/fizz_buzz/logic"
)

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
	}
}

func print(number int) {
	fmt.Println(logic.FizzBuzzGenerator(number))
}
