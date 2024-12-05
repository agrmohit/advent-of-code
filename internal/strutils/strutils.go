// Package strutils contains commonly used string functions
package strutils

// Reverse takes a string and returns the reversed string
//
// Copied from: https://groups.google.com/g/golang-nuts/c/oPuBaYJ17t4/m/PCmhdAyrNVkJ,
// Russ Cox, golang-nuts mailing list
func Reverse(input string) string {
	// Get Unicode code points.
	n := 0
	characterRune := make([]rune, len(input))

	for _, r := range input {
		characterRune[n] = r
		n++
	}

	characterRune = characterRune[0:n]

	// Reverse
	for i := 0; i < n/2; i++ {
		characterRune[i], characterRune[n-1-i] = characterRune[n-1-i], characterRune[i]
	}

	// Convert back to UTF-8.
	output := string(characterRune)
	return output
}
