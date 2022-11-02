package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arraySize int

func main() {

	// Taking the input from the user and storing it in the variable arraySize.
	fmt.Scan(&arraySize)

	// Creating a slice of type int with the size of arraySize.
	array := make([]int, arraySize)

	// Creating a new scanner object which will read from the standard input.
	arrayElements := bufio.NewScanner(os.Stdin)
	// Scanning the input from the standard input.
	arrayElements.Scan()

	// Storing the input from the standard input in the variable data.
	data := arrayElements.Text()
	// Splitting the string data into a slice of strings.
	newData := strings.Split(data, " ")

	// Storing the length of the array in the variable size.
	size := len(array)

	for _, k := range newData {
		array[size-1], _ = strconv.Atoi(k)
		size--
	}

	for _, k := range array {
		fmt.Printf("%d ", k)
	}

}
