package main

import (
	"fmt"
)

func main() {
	numbers := []int{4, 1, 3, 7, 8, 9, 4, 7, 1, 1, 7, 5, 5, 9, 0, 4}
	fmt.Println(validateCardNumber(numbers))
}

func validateCardNumber(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			numbers[i] = numbers[i] * 2
		}
		if numbers[i] > 9 {
			numbers[i] = sumFromNumber(numbers[i])
		}
	}

	sumOfDigit := 0
	for i := 0; i < len(numbers); i++ {
		sumOfDigit += numbers[i]
	}
	return (sumOfDigit % 10) == 0
}

func sumFromNumber(number int) int {
	sum := 0
	for {
		sum += number % 10
		number = number / 10
		if number <= 0 {
			break
		}
	}
	return sum
}
