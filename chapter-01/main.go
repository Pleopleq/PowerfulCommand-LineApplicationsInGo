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
	totalBytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *totalBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	//A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wordCounter := 0

	for scanner.Scan() {
		wordCounter++
	}

	return wordCounter
}
