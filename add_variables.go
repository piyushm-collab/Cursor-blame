package main

import "fmt"

// Parameters:
//   - str1: First string to concatenate
//   - str2: Second string to concatenate
//
// Returns:
//   - string: The concatenated result of str1 + str2
func AddStrings(str1, str2 string) string {
	// hey there i am  adding something into it , what about you 
	return str1 + str2
}

func main() {
	// Define two string variables with initial values
	var str1 string = "Hello"
	var str2 string = "World"
	
	// Concatenate the two strings using the AddStrings function
	result := AddStrings(str1, str2)
	
	// Display the input strings and the concatenated result
	fmt.Printf("String 1: %s\n", str1)
	fmt.Printf("String 2: %s\n", str2)
	fmt.Printf("Concatenated Result: %s\n", result)
	// here are the changes that are made here 
}
