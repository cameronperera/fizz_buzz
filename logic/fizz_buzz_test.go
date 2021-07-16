package logic

import (
	"testing"
)

func TestFizzBuzzGeneratorWhenInputIsNotDivisibleByThreeOrFive(t *testing.T) {
	stringReturned := FizzBuzzGenerator(1)

	expectedString := "1"
	if stringReturned != expectedString {
		t.Errorf("It should return %d since it is not divisible by 3 or 5", 1)
	}
}

func TestFizzWhenInputIsDivisibleByThree(t *testing.T) {
	stringReturned := FizzBuzzGenerator(3)

	expectedString := "Fizz"
	if stringReturned != expectedString {
		t.Errorf("It should print Fizz for number %d since it is divisible by 3", 3)
	}
}

func TestBuzzWhenInputIsDivisibleByFive(t *testing.T) {
	stringReturned := FizzBuzzGenerator(5)

	expectedString := "Buzz"
	if stringReturned != expectedString {
		t.Errorf("It should print Buzz for number %d since it is divisible by 5", 5)
	}
}

func TestFizzBuzzWhenInputIsDivisibleByThreeAndFive(t *testing.T) {
	stringReturned := FizzBuzzGenerator(15)

	expectedString := "FizzBuzz"
	if stringReturned != expectedString {
		t.Errorf("It should print FizzBuzz for number %d since it is divisible by 3 & 5", 15)
	}
}
