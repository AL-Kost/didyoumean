package main

import (
	"flag"
	"fmt"
	"github.com/AL-Kost/didyoumean"
	"os"
	"strings"
)

// contains checks if a slice contains a specific string
// It's used to check if a flag is a valid, predefined flag.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	// Define flags for the CLI tool
	wordPtr := flag.String("word", "", "The word for which you want to find suggestions.")
	helpPtr := flag.Bool("help", false, "Display help information.")

	// A predefined list of valid flag names
	validFlags := []string{"word", "help"}

	// Manually parse command-line arguments before calling flag.Parse()
	for _, arg := range os.Args[1:] { // os.Args[0] is the program name, so we skip it
		// Check if the argument is a flag (starts with a hyphen)
		if strings.HasPrefix(arg, "-") {
			// Trim the hyphen to get the flag name
			flagName := strings.TrimPrefix(arg, "-")
			// Check if the provided flag is not in the list of valid flags
			if !contains(validFlags, flagName) {
				// Use the didyoumean package to suggest the closest match to the invalid flag
				suggestion := didyoumean.Suggest(flagName, validFlags, 1, true)
				fmt.Printf("Unknown flag '-%s' provided.\n", flagName)
				if suggestion != "" {
					fmt.Printf("Did you mean '-%s'?\n", suggestion)
				}
				// Display usage information and exit with a non-zero status code
				flag.Usage()
				os.Exit(2)
			}
		}
	}

	// After checking for unknown flags, we can safely parse the known flags
	flag.Parse()

	// If the help flag is provided, print out usage information
	if *helpPtr {
		flag.Usage()
		return
	}

	// Check if the user has provided the word flag
	if *wordPtr == "" {
		fmt.Println("Please provide a word to find suggestions for using the -word flag.")
		os.Exit(2)
	}

	// Predefined list of candidate strings to match against
	candidates := []string{"apple", "application", "apology", "apply", "happy"}

	// Use the didyoumean package to suggest the closest match to the input word
	closestMatch := didyoumean.Suggest(*wordPtr, candidates, 2, true)

	// Output the result
	if closestMatch != "" {
		fmt.Printf("Did you mean '%s'?\n", closestMatch)
	} else {
		fmt.Println("No close match found.")
	}
}
