package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Reading the input from the user and storing it in the variable text.
	inputscan := bufio.NewReader(os.Stdin)
	text, _ := inputscan.ReadString('\n')

	// Printing the text that the user entered.
	fmt.Println("Hello, World.")
	fmt.Println(text)
}
