package enigma

type rotor struct {
	sequence         [numberOfLetters]int
	invertedSequence [numberOfLetters]int
	turnover         []int
	offset           int
	ring             int
}

func newRotor(mapping, turnovers string) rotor {
	r := rotor{}

	r.turnover = make([]int, len(turnovers))
	for i := range turnovers {
		r.turnover[i] = char(turnovers[i])
	}

	for i, letter := range mapping {
		index := char(byte(letter))
		r.sequence[i] = index
		r.invertedSequence[index] = i
	}

	return r
}

func (r *rotor) setConfiguration(offset, ring int) {
	r.offset = offset
	r.ring = ring
}

func (r *rotor) move() {
	r.offset = (r.offset + 1) % numberOfLetters
}

func (r rotor) isTurnOver() bool {
	for _, turnover := range r.turnover {
		if r.offset == turnover {
			return true
		}
	}

	return false
}

func (r rotor) step(letter int, invert bool) int {
	letter = (letter - r.ring + r.offset + numberOfLetters) % numberOfLetters

	if invert {
		letter = r.invertedSequence[letter]
	} else {
		letter = r.sequence[letter]
	}

	letter = (letter + r.ring - r.offset + numberOfLetters) % numberOfLetters

	return letter
}
