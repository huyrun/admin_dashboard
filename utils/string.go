package utils

import "unicode"

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s // Return empty string if input is empty
	}

	runes := []rune(s) // Convert to rune slice to handle Unicode characters
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
