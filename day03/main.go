package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	if N%2 != 0 {
		fmt.Println("Weird")
	}

	if N%2 == 0 {
		for i := 2; i < 6; i++ {
			if int(N) == i {
				fmt.Println("Not Weird")
			}
		}
	}

	if N%2 == 0 {
		for i := 6; i < 21; i++ {
			if int(N) == i {
				fmt.Println("Weird")
			}
		}
	}

	if N%2 == 0 && N > 20 {
		fmt.Println("Not Weird")
	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
