package enigma

type plugboard [numberOfLetters]int

func newPlugboard(pairs []string) plugboard {
	p := plugboard{}
	for i := 0; i < numberOfLetters; i++ {
		p[i] = i
	}

	for _, pair := range pairs {
		if len(pair) > 0 {
			var intFirst = char(pair[0])
			var intSecond = char(pair[1])
			p[intFirst] = intSecond
			p[intSecond] = intFirst
		}
	}

	return p
}
