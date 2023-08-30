package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	//A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	wordCounter := 0

	for scanner.Scan() {
		wordCounter++
	}

	return wordCounter
}
