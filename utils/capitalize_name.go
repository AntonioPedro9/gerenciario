package utils

import "strings"

func CapitalizeName(name string) string {
	words := strings.Fields(name)
	capitalizedWords := make([]string, len(words))

	for i, word := range words {
		lowercaseWord := strings.ToLower(word)
		firstLetter := strings.ToUpper(string(lowercaseWord[0]))
		capitalizedWords[i] = firstLetter + lowercaseWord[1:]
	}

	return strings.Join(capitalizedWords, " ")
}
