# Enigma machine in Golang

Simple enigma implementation in pure Golang

## Features

Historical configuration of enigma:

### Rotors:

- `I`:     `EKMFLGDQVZNTOWYHXUSPAIBRCJ`, turnover: `Q`
- `II`:    `AJDKSIRUXBLHWTMCQGZNPYFVOE`, turnover: `E`
- `III`:   `BDFHJLCPRTXVZNYEIWGAKMUSQO`, turnover: `V`
- `IV`:    `ESOVPZJAYQUIRHXLNFTGKDCMWB`, turnover: `J`
- `V`:     `VZBRGITYUPSDNHLXAWMJQOFECK`, turnover: `Z`
- `VI`:    `JPGVOUMFYQBENHZRDKASXLICTW`, turnover: `ZM`
- `VII`:   `NZJHGRCXMYSWBOUFAIVLPEKQDT`, turnover: `ZM`
- `VIII`:  `FKQHTLXOCBJSPDZRAMEWNIUYGV`, turnover: `ZM`
- `Beta`:  `LEYJVCNIXWPBQMDRTAKZGFUHOS`, turnover: -
- `Gamma`: `FSOKANUERHMBTIYCWLQPZXVGJD`, turnover: -

### Reflectors:

- `A`:      `EJMZALYXVBWFCRQUONTSPIKHGD`
- `B`:      `YRUHQSLDPXNGOKMIEBFZCWVJAT`
- `C`:      `FVPJIAOYEDRZXWGCTKUQSBNMHL`
- `B-thin`: `ENKQAUYWJICOPBLMDXZVFTHRGS`
- `C-thin`: `RDOBJNTKVEHMLFCWZAXGYIPSUQ`