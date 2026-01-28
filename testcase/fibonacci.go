package main

import "fmt"

// FibonacciSeries generates and prints the Fibonacci series up to n terms
// The Fibonacci series is a sequence where each number is the sum of the two preceding ones
func FibonacciSeries(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
	// here is the program for the fibonacci series
}
//DUmmy comment
// FibonacciRecursive calculates the nth Fibonacci number using recursion
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func main() {
	// Number of terms to generate
	numTerms := 10

	fmt.Println("Fibonacci Series (Iterative):")
	fibSeries := FibonacciSeries(numTerms)
	for i, val := range fibSeries {
		fmt.Printf("Term %d: %d\n", i+1, val)
	}

	fmt.Println("\nFibonacci Series (Recursive):")
	for i := 0; i < numTerms; i++ {
		fmt.Printf("Term %d: %d\n", i+1, FibonacciRecursive(i))
	}

	fmt.Println("\nFirst 15 Fibonacci numbers:")
	fibSeries15 := FibonacciSeries(15)
	for i, val := range fibSeries15 {
		fmt.Printf("%d ", val)
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
