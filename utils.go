package enigma

import (
	"regexp"
	"strings"
)

const (
	separator        = ","
	numberOfLetters  = 26
	firstLetter      = 'A'
	numberOfRotors   = 3
	spaceReplacement = "X"
)

func char(char byte) int {
	return int(char - firstLetter)
}

func index(index int) byte {
	return byte(firstLetter + index)
}

func cleanUp(plaintext string) string {
	plaintext = strings.TrimSpace(plaintext)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.ReplaceAll(plaintext, " ", spaceReplacement)
	plaintext = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(plaintext, "")

	return plaintext
}
