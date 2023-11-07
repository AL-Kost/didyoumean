package main

import (
	"fmt"
	"github.com/AL-Kost/didyoumean"
)

func main() {
	// Example usage of the didyoumean package

	// Input string we're trying to find a match for
	input := "appl"

	// Candidate strings to match against the input
	candidates := []string{"apple", "application", "apology", "apply", "happy"}

	// Threshold for Levenshtein distance
	// If the closest match has a distance greater than the threshold, no match will be suggested
	threshold := 2

	// Set to true if the comparison should be case-insensitive
	caseInsensitive := true

	// Call the Suggest function from the didyoumean package
	closestMatch := didyoumean.Suggest(input, candidates, threshold, caseInsensitive)

	// Output the closest match to the console
	if closestMatch != "" {
		fmt.Printf("Did you mean '%s'?\n", closestMatch)
	} else {
		fmt.Println("No close match found.")
	}
}
