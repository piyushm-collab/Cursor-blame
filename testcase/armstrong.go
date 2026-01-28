package main

import (
	"fmt"
	"math"
)

// isArmstrong checks if a number is an Armstrong number
// An Armstrong number is a number that is equal to the sum of its digits
// each raised to the power of the number of digits
func isArmstrong(num int) bool {
	if num < 0 {
		return false
	}

	original := num
	sum := 0
	digits := countDigits(num)

	for num > 0 {
		digit := num % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		num /= 10
	}

	return sum == original
}

// countDigits returns the number of digits in a number
func countDigits(num int) int {
	if num == 0 {
		return 1
	}

	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count
}

// findArmstrongNumbers finds all Armstrong numbers in a given range
func findArmstrongNumbers(start, end int) []int {
	var armstrongNumbers []int

	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongNumbers = append(armstrongNumbers, i)
		}
	}

	return armstrongNumbers
}

func main() {
	// Test individual numbers
	testNumbers := []int{0, 1, 153, 370, 371, 407, 1634, 8208, 9474, 100, 200}

	fmt.Println("Testing individual numbers:")
	for _, num := range testNumbers {
		if isArmstrong(num) {
			fmt.Printf("%d is an Armstrong number\n", num)
		} else {
			// here is the program
			fmt.Printf("%d is not an Armstrong number\n", num)
		}
	}

	// Find all Armstrong numbers in a range
	fmt.Println("\nArmstrong numbers from 1 to 1000:")
	armstrongNums := findArmstrongNumbers(1, 1000)
	fmt.Println(armstrongNums)

	fmt.Println("\nArmstrong numbers from 1000 to 10000:")
	armstrongNums = findArmstrongNumbers(1000, 10000)
	fmt.Println(armstrongNums)
}
