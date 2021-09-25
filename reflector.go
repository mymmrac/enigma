package enigma

type reflector struct {
	ID       string
	Sequence [26]int
}

func newReflector(mapping string, id string) *reflector {
	var seq [26]int
	for i, value := range mapping {
		seq[i] = char(byte(value))
	}
	return &reflector{id, seq}
}

type reflectors []reflector

func (refs *reflectors) GetByID(id string) *reflector {
	for _, ref := range *refs {
		if ref.ID == id {
			return &ref
		}
	}
	return nil
}
