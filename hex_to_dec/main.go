package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Enter hex number or 'stop' to exit:")
	
	var input string
	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "stop" {
			break
		}

		i := new(big.Int)

		// _, ok := i.SetString(processHex(input), 16) // 0xabc123

		// if !ok {
		// 	fmt.Println("Failed!")
		// } else {
		// 	fmt.Println(i)
		// }

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Failed!")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.Trim(hexStr, "0x")
}