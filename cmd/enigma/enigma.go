package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mymmrac/enigma"
)

func main() {
	rotorsOptions := flag.String("rotors", "I,II,III", "Rotors configuration")
	positionOptions := flag.String("positions", "A,A,A", "Positions of rotors")
	ringsOptions := flag.String("rings", "1,1,1", "Rotor rings offset")
	reflectorOption := flag.String("reflector", "A", "Reflector")
	plugboardOptions := flag.String("plugboard", "", "Plugboard configuration")

	verbose := flag.Bool("verbose", false, "Verbose output")

	configs := flag.Bool("configs", false, "Show available configs")

	flag.Parse()

	if *configs {
		enigma.PrintConfigs()
		return
	}

	e, err := enigma.NewEnigmaFromConfig(*rotorsOptions, *positionOptions, *ringsOptions, *reflectorOption, *plugboardOptions)
	if err != nil {
		fmt.Println("Invalid configuration:", err)
		os.Exit(1)
	}

	text := strings.Join(flag.Args(), " ")
	if len(text) == 0 {
		reader := bufio.NewReader(os.Stdin)
		text, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while getting text:", err)
			os.Exit(1)
		}
	}

	encoded := e.EncodeString(text)

	if *verbose {
		fmt.Printf("Original text:\n%s\n", text)
		fmt.Printf("Cipher text:\n%s\n", encoded)
		return
	}

	fmt.Println(encoded)
}
