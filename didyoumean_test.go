package didyoumean

import "testing"

func TestSuggest(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		candidates      []string
		threshold       int
		caseInsensitive bool
		want            string
	}{
		{"Exact match", "apple", []string{"apple", "banana", "apricot"}, 0, false, "apple"},
		{"Case insensitive match", "Apple", []string{"apple", "banana", "apricot"}, 0, true, "apple"},
		{"No match within threshold", "apple", []string{"banana", "grape", "orange"}, 2, false, ""},
		{"Match within threshold", "aple", []string{"apple", "apricot", "banana"}, 1, false, "apple"},
		{"Case insensitive no match", "Apple", []string{"banana", "grape", "orange"}, 2, true, ""},
		{"Empty input and candidates", "", []string{""}, 0, false, ""},
		{"No candidates", "apple", []string{}, 0, false, ""},
		{"Input longer than candidate", "apple", []string{"app"}, 10, false, "app"},
		{"Candidate longer than input", "app", []string{"apple"}, 10, false, "apple"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Suggest(tt.input, tt.candidates, tt.threshold, tt.caseInsensitive); got != tt.want {
				t.Errorf("Suggest(%q, %#v, %d, %t) = %q; want %q", tt.input, tt.candidates, tt.threshold, tt.caseInsensitive, got, tt.want)
			}
		})
	}
}

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want int
	}{
		{"Same strings", "distance", "distance", 0},
		{"One character add", "distance", "distances", 1},
		{"One character remove", "distance", "istance", 1},
		{"One character replace", "distance", "distanse", 1},
		{"Completely different", "distance", "apple", 7},
		{"First string empty", "", "distance", 8},
		{"Second string empty", "distance", "", 8},
		{"Both strings empty", "", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenshteinDistance(tt.s1, tt.s2); got != tt.want {
				t.Errorf("levenshteinDistance(%q, %q) = %d; want %d", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestMinSet(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"Single element", []int{1}, 1},
		{"All elements same", []int{2, 2, 2}, 2},
		{"Mixed values", []int{1, 3, 2}, 1},
		{"Negative values", []int{-1, -3, -2}, -3},
		{"Mixed negative and positive", []int{-1, 1, 0}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSet(tt.values...); got != tt.want {
				t.Errorf("minSet(%v) = %d; want %d", tt.values, got, tt.want)
			}
		})
	}
}
