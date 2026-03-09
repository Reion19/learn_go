package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

// var cipherMode = flag.Bool("cipher", true, "Enable cipher mode. This is default mode.")
// var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")
var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character.")

func main() {
	flag.Parse()

	// if *cipherMode && *decipherMode {
	// 	fmt.Println("Please specify only one mode at a time.")
	// 	os.Exit(1)
	// }

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter your text to cipher: ")

		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ecnrypted text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		cipheredtext := getUserInput("Enter your ciphered data to decipher: ")

		decipheredText, err := cipherer.Decipher(cipheredtext, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ecnrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Println("Please specify only one mode at a time.")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided! Exiting now...")
		os.Exit(1)
	}
}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading the entered  text! Please try again.")
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}
