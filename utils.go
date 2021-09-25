package enigma

import (
	"regexp"
	"strings"
)

func char(char byte) int {
	return int(char - 'A')
}

func index(index int) byte {
	return byte('A' + index)
}

func cleanUp(plaintext string) string {
	plaintext = strings.TrimSpace(plaintext)
	plaintext = strings.ToUpper(plaintext)
	plaintext = strings.Replace(plaintext, " ", "X", -1)
	plaintext = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(plaintext, "")

	return plaintext
}
