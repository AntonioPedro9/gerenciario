package utils

import "strings"

func FormatName(name string) (string, error) {
	if len(name) < 2 {
		return "", InvalidNameError
	}

	words := strings.Fields(name)
	capitalizedWords := make([]string, len(words))

	for i, word := range words {
		lowercaseWord := strings.ToLower(word)
		firstLetter := strings.ToUpper(string(lowercaseWord[0]))
		capitalizedWords[i] = firstLetter + lowercaseWord[1:]
	}

	capitalizedName := strings.Join(capitalizedWords, " ")

	return capitalizedName, nil
}
