package enigma

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

const separator = ","

type Enigma struct {
	reflector reflector
	plugboard plugboard
	rotors    []*rotor
}

type RotorConfig struct {
	ID    string
	Start byte
	Ring  int
}

func NewEnigma(rotorConfiguration []RotorConfig, refID string, plugs []string) *Enigma {
	rotors := make([]*rotor, len(rotorConfiguration))
	for i, configuration := range rotorConfiguration {
		rotors[i] = historicRotors.GetByID(configuration.ID)
		rotors[i].Offset = char(configuration.Start)
		rotors[i].Ring = configuration.Ring - 1
	}

	return &Enigma{
		reflector: *historicReflectors.GetByID(refID),
		plugboard: *newPlugboard(plugs),
		rotors:    rotors,
	}
}

func NewEnigmaFromConfig(rotors, positions, rings, reflector, plugboard string) (*Enigma, error) {
	if err := validateFormat(rotors, positions, rings, reflector, plugboard); err != nil {
		return nil, err
	}

	var rotorsConfigs []RotorConfig

	rotorIDs := strings.Split(rotors, separator)
	positionChars := strings.Split(positions, separator)
	ringsNumbers := strings.Split(rings, separator)

	if len(rotorIDs) != len(positionChars) || len(rotorIDs) != len(ringsNumbers) {
		return nil, errors.New("number of rotors, positions and rings must be the same")
	}

	for i := 0; i < len(rotorIDs); i++ {
		ring, _ := strconv.Atoi(ringsNumbers[i])

		rotorsConfigs = append(rotorsConfigs, RotorConfig{
			ID:    rotorIDs[i],
			Start: positionChars[i][0],
			Ring:  ring,
		})
	}

	plugboardPairs := strings.Split(plugboard, separator)

	return NewEnigma(rotorsConfigs, reflector, plugboardPairs), nil
}

func (e *Enigma) moveRotors() {
	var (
		rotorLen            = len(e.rotors)
		farRight            = e.rotors[rotorLen-1]
		farRightTurnover    = farRight.ShouldTurnOver()
		secondRight         = e.rotors[rotorLen-2]
		secondRightTurnover = secondRight.ShouldTurnOver()
		thirdRight          = e.rotors[rotorLen-3]
	)

	if secondRightTurnover {
		if !farRightTurnover {
			secondRight.move(1)
		}
		thirdRight.move(1)
	}

	if farRightTurnover {
		secondRight.move(1)
	}

	farRight.move(1)
}

func (e *Enigma) EncodeChar(letter byte) byte {
	e.moveRotors()

	letterIndex := char(letter)
	letterIndex = e.plugboard[letterIndex]

	for i := len(e.rotors) - 1; i >= 0; i-- {
		letterIndex = e.rotors[i].Step(letterIndex, false)
	}

	letterIndex = e.reflector.Sequence[letterIndex]

	for i := 0; i < len(e.rotors); i++ {
		letterIndex = e.rotors[i].Step(letterIndex, true)
	}

	letterIndex = e.plugboard[letterIndex]
	letter = index(letterIndex)

	return letter
}

func (e *Enigma) EncodeString(text string) string {
	text = cleanUp(text)

	var result bytes.Buffer
	for i := range text {
		result.WriteByte(e.EncodeChar(text[i]))
	}

	return result.String()
}
