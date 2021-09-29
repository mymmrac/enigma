package enigma

import "fmt"

var historicalRotors = map[string]rotor{
	"I":     newRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q"),
	"II":    newRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "E"),
	"III":   newRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "V"),
	"IV":    newRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J"),
	"V":     newRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "Z"),
	"VI":    newRotor("JPGVOUMFYQBENHZRDKASXLICTW", "ZM"),
	"VII":   newRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "ZM"),
	"VIII":  newRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "ZM"),
	"Beta":  newRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", ""),
	"Gamma": newRotor("FSOKANUERHMBTIYCWLQPZXVGJD", ""),
}

var historicalReflectors = map[string]reflector{
	"A":      newReflector("EJMZALYXVBWFCRQUONTSPIKHGD"),
	"B":      newReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT"),
	"C":      newReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL"),
	"B-thin": newReflector("ENKQAUYWJICOPBLMDXZVFTHRGS"),
	"C-thin": newReflector("RDOBJNTKVEHMLFCWZAXGYIPSUQ"),
}

// PrintConfigs prints available configuration options of enigma
func PrintConfigs() {
	fmt.Println("Rotors:")
	for name, rotor := range historicalRotors {
		mapping := ""
		for _, ch := range rotor.sequence {
			mapping += string(rune(firstLetter + ch))
		}
		turnovers := ""
		for _, ch := range rotor.turnover {
			turnovers += string(rune('A' + ch))
		}
		fmt.Printf("Name: %q, Mapping: %q, Turnovers: %q\n", name, mapping, turnovers)
	}
	fmt.Println()

	fmt.Println("Positions: A-Z")
	fmt.Println()

	fmt.Println("Rings: 1-26")
	fmt.Println()

	fmt.Println("Reflectors:")
	for name, reflector := range historicalReflectors {
		mapping := ""
		for _, ch := range reflector {
			mapping += string(rune('A' + ch))
		}
		fmt.Printf("Name: %q, Mapping: %q\n", name, mapping)
	}
	fmt.Println()

	fmt.Println("Plugboard: any sequence of letter pairs")
}
