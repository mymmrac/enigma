package enigma

import "fmt"

var historicRotors = rotors{
	*newRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "I", "Q"),
	*newRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "II", "E"),
	*newRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "III", "V"),
	*newRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "IV", "J"),
	*newRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "V", "Z"),
	*newRotor("JPGVOUMFYQBENHZRDKASXLICTW", "VI", "ZM"),
	*newRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "VII", "ZM"),
	*newRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "VIII", "ZM"),
	*newRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", "Beta", ""),
	*newRotor("FSOKANUERHMBTIYCWLQPZXVGJD", "Gamma", ""),
}

var historicReflectors = reflectors{
	*newReflector("EJMZALYXVBWFCRQUONTSPIKHGD", "A"),
	*newReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT", "B"),
	*newReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL", "C"),
	*newReflector("ENKQAUYWJICOPBLMDXZVFTHRGS", "B-thin"),
	*newReflector("RDOBJNTKVEHMLFCWZAXGYIPSUQ", "C-thin"),
}

func PrintConfigs() {
	fmt.Println("Rotors:")
	for _, rotor := range historicRotors {
		mapping := ""
		for _, ch := range rotor.StraightSeq {
			mapping += string(rune('A' + ch))
		}
		turnovers := ""
		for _, ch := range rotor.Turnover {
			turnovers += string(rune('A' + ch))
		}
		fmt.Printf("Name: %q, Mapping: %q, Turnovers: %q\n", rotor.ID, mapping, turnovers)
	}
	fmt.Println()

	fmt.Println("Positions: A-Z")
	fmt.Println()

	fmt.Println("Rings: 1-26")
	fmt.Println()

	fmt.Println("Reflectors:")
	for _, reflector := range historicReflectors {
		mapping := ""
		for _, ch := range reflector.Sequence {
			mapping += string(rune('A' + ch))
		}
		fmt.Printf("Name: %q, Mapping: %q\n", reflector.ID, mapping)
	}
	fmt.Println()

	fmt.Println("Plugboard: any sequence of letter pairs")
}
