package enigma

type rotor struct {
	ID          string
	StraightSeq [26]int
	ReverseSeq  [26]int
	Turnover    []int

	Offset int
	Ring   int
}

func newRotor(mapping string, id string, turnovers string) *rotor {
	r := &rotor{ID: id, Offset: 0, Ring: 0}
	r.Turnover = make([]int, len(turnovers))
	for i := range turnovers {
		r.Turnover[i] = char(turnovers[i])
	}

	for i, letter := range mapping {
		index := char(byte(letter))
		r.StraightSeq[i] = index
		r.ReverseSeq[index] = i
	}

	return r
}

func (r *rotor) move(offset int) {
	r.Offset = (r.Offset + offset) % 26
}

func (r *rotor) ShouldTurnOver() bool {
	for _, turnover := range r.Turnover {
		if r.Offset == turnover {
			return true
		}
	}

	return false
}

func (r *rotor) Step(letter int, invert bool) int {
	letter = (letter - r.Ring + r.Offset + 26) % 26

	if invert {
		letter = r.ReverseSeq[letter]
	} else {
		letter = r.StraightSeq[letter]
	}

	letter = (letter + r.Ring - r.Offset + 26) % 26

	return letter
}

type rotors []rotor

func (rs *rotors) GetByID(id string) *rotor {
	for _, rotor := range *rs {
		if rotor.ID == id {
			return &rotor
		}
	}

	return nil
}
