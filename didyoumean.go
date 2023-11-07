package didyoumean

import "strings"

// Suggest finds the closest match for a given input string from a list of candidate strings.
// It uses the Levenshtein distance algorithm to determine the closest match.
// If no candidate is within the specified threshold, an empty string is returned.
// The comparison can be case-insensitive if the caseInsensitive flag is set.
//
// Parameters:
// - input: the input string to match.
// - candidates: a slice of candidate strings to compare against the input.
// - threshold: the maximum Levenshtein distance considered for a match to be close enough.
// - caseInsensitive: if set to true, the matching will be case-insensitive.
//
// Returns:
// - The closest match within the threshold, or an empty string if no match is close enough.
func Suggest(input string, candidates []string, threshold int, caseInsensitive bool) string {
	if caseInsensitive {
		input = strings.ToLower(input)
	}

	minDistance := int(^uint(0) >> 1) // Initialize with the largest possible int value
	closest := ""

	for _, candidate := range candidates {
		candidateToCompare := candidate
		if caseInsensitive {
			candidateToCompare = strings.ToLower(candidate)
		}

		distance := levenshteinDistance(input, candidateToCompare)
		if distance < minDistance {
			minDistance = distance
			closest = candidate
		}
	}

	if minDistance <= threshold {
		return closest
	}

	return "" // No match is close enough within the given threshold
}

// levenshteinDistance calculates the Levenshtein distance which is a measure of
// the difference between two sequences. In terms of strings, it is the minimum
// number of single-character edits required to change one word into the other.
// This implementation uses dynamic programming to efficiently compute the distances
// by storing the results of subproblems in a 2D slice. The function is case-sensitive
// and operates on runes to correctly handle Unicode characters, which may be
// multiple bytes long in string representation.
//
// Parameters:
// - s1: the first string.
// - s2: the second string.
//
// Returns:
// - The Levenshtein distance between the two strings.
func levenshteinDistance(s1, s2 string) int {
	// conversion to rune slices is necessary to support Unicode characters
	a, b := []rune(s1), []rune(s2)
	// distances is a slice of slices, initialized to store the distances
	// where distances[i][j] will hold the Levenshtein distance between
	// the first i characters of s1 and the first j characters of s2.
	distances := make([][]int, len(a)+1)
	for i := range distances {
		distances[i] = make([]int, len(b)+1)
	}
	// initialize the first row and column with distances corresponding
	// to transforming a string into an empty string by dropping all characters
	for i := 0; i <= len(a); i++ {
		distances[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		distances[0][j] = j
	}
	// calculate distances using the previously defined recurrence relation
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1 // cost is 1 when characters do not match
			}
			// set the distance for the current subproblem by considering
			// the minimum cost among deletion, insertion, and substitution
			distances[i][j] = minSet(
				distances[i-1][j]+1,      // deletion
				distances[i][j-1]+1,      // insertion
				distances[i-1][j-1]+cost, // substitution
			)
		}
	}
	// the bottom-right cell contains the Levenshtein distance between s1 and s2
	return distances[len(a)][len(b)]
}

// minSet returns the smallest integer from a set of integers.
// It is a helper function used within levenshteinDistance to determine
// the minimum cost between deletion, insertion, and substitution operations.
//
// Parameters:
// - values: a variable number of int arguments to find the minimum value of.
//
// Returns:
// - The smallest integer among the provided values.
func minSet(values ...int) int {
	minimum := values[0]
	for _, val := range values[1:] {
		if val < minimum {
			minimum = val
		}
	}
	return minimum
}
