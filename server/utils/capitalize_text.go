package utils

import "strings"

func CapitalizeText(text string) string {
	words := strings.Fields(text)
	capitalizedWords := make([]string, len(words))

	for i, word := range words {
		lowercaseWord := strings.ToLower(word)
		firstLetter := strings.ToUpper(string(lowercaseWord[0]))
		capitalizedWords[i] = firstLetter + lowercaseWord[1:]
	}

	return strings.Join(capitalizedWords, " ")
}
