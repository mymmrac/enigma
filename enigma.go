package enigma

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Enigma represents enigma machine
type Enigma struct {
	reflector reflector
	plugboard plugboard
	rotors    []rotor
}

// RotorConfig represents rotor configuration
type RotorConfig struct {
	Name  string
	Start byte
	Ring  int
}

// NewEnigma creates new enigma
func NewEnigma(rotorConfiguration []RotorConfig, reflectorName string, plugs []string) (*Enigma, error) {
	ok := false
	rotors := make([]rotor, len(rotorConfiguration))
	for i, configuration := range rotorConfiguration {
		rotors[i], ok = historicalRotors[configuration.Name]
		if !ok {
			return nil, errors.New("rotor not found")
		}

		rotors[i].setConfiguration(char(configuration.Start), configuration.Ring-1)
	}

	selectedReflector, ok := historicalReflectors[reflectorName]
	if !ok {
		return nil, errors.New("reflector not found")
	}

	return &Enigma{
		reflector: selectedReflector,
		plugboard: newPlugboard(plugs),
		rotors:    rotors,
	}, nil
}

// NewEnigmaParse creates new enigma from configuration
func NewEnigmaParse(rotors, positions, rings, reflector, plugboard string) (*Enigma, error) {
	if err := validateFormat(rotors, positions, rings, reflector, plugboard); err != nil {
		return nil, err
	}

	var rotorsConfigs []RotorConfig

	rotorNames := strings.Split(rotors, separator)
	positionChars := strings.Split(positions, separator)
	ringsNumbers := strings.Split(rings, separator)

	if len(rotorNames) != numberOfRotors ||
		len(positionChars) != numberOfRotors ||
		len(ringsNumbers) != numberOfRotors {
		return nil, errors.New("number of rotors, positions and rings must be the same")
	}

	for i := 0; i < len(rotorNames); i++ {
		ring, _ := strconv.Atoi(ringsNumbers[i])

		rotorsConfigs = append(rotorsConfigs, RotorConfig{
			Name:  rotorNames[i],
			Start: positionChars[i][0],
			Ring:  ring,
		})
	}

	plugboardPairs := strings.Split(plugboard, separator)

	return NewEnigma(rotorsConfigs, reflector, plugboardPairs)
}

func (e *Enigma) moveRotors() {
	if e.rotors[1].isTurnOver() {
		if !e.rotors[2].isTurnOver() {
			e.rotors[1].move()
		}
		e.rotors[0].move()
	}

	if e.rotors[2].isTurnOver() {
		e.rotors[1].move()
	}

	e.rotors[2].move()
}

func (e *Enigma) EncodeChar(letter byte) byte {
	e.moveRotors()

	letterIndex := char(letter)
	letterIndex = e.plugboard[letterIndex]

	for i := len(e.rotors) - 1; i >= 0; i-- {
		letterIndex = e.rotors[i].step(letterIndex, false)
	}

	letterIndex = e.reflector[letterIndex]

	for i := 0; i < len(e.rotors); i++ {
		letterIndex = e.rotors[i].step(letterIndex, true)
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
