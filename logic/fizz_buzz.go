package logic

import (
	"strconv"
)

func FizzBuzzGenerator(number int) string {
	if numberIsDisivibleByThree(number) && numberIsDisivibleByFive(number) {
		return ("FizzBuzz")
	}
	if numberIsDisivibleByFive(number) {
		return ("Buzz")
	}
	if numberIsDisivibleByThree(number) {
		return ("Fizz")
	}

	return (strconv.Itoa(number))
}

func numberIsDisivibleByThree(number int) bool {
	return number%3 == 0
}

func numberIsDisivibleByFive(number int) bool {
	return number%5 == 0
}
