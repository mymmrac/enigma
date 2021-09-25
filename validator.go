package enigma

import (
	"errors"
	"regexp"
)

var rotorsRegex = regexp.MustCompile(`^(:?\w+,)*?(:?\w+)$`)
var positionsRegex = regexp.MustCompile(`^(:?[A-Z]+,)*?(:?[A-Z]+)$`)
var ringsRegex = regexp.MustCompile(`^(:?[1-9][0-9]*,)*?(:?[1-9][0-9]*)$`)
var reflectorRegexp = regexp.MustCompile(`^\w+$`)
var plugboardRegexp = regexp.MustCompile(`^(:?[A-Z]{2},)*?(:?[A-Z]{2})?$`)

func validateFormat(rotors, positions, rings, reflector, plugboard string) error {
	if !rotorsRegex.MatchString(rotors) {
		return errors.New("rotors format")
	}

	if !positionsRegex.MatchString(positions) {
		return errors.New("positions format")
	}

	if !ringsRegex.MatchString(rings) {
		return errors.New("positions format")
	}

	if !reflectorRegexp.MatchString(reflector) {
		return errors.New("reflector format")
	}

	if !plugboardRegexp.MatchString(plugboard) {
		return errors.New("reflector format")
	}

	return nil
}
