package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	//A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wordCounter := 0

	for scanner.Scan() {
		wordCounter++
	}

	return wordCounter
}
