package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var (
		i2 uint64
		d2 float64
		s2 string
	)

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&i2)
	fmt.Scan(&d2)
	scanner.Scan()
	s2 = scanner.Text()
	fmt.Println(i + i2)
	fmt.Printf("%.1f\n", d+d2)
	fmt.Println(s + s2)
}
