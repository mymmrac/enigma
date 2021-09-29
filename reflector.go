package enigma

type reflector [numberOfLetters]int

func newReflector(mapping string) reflector {
	r := reflector{}
	for i, value := range mapping {
		r[i] = char(byte(value))
	}

	return r
}
