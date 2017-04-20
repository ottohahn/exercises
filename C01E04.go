package main

// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	infile := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, infile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, infile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, infile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, infile map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		infile[input.Text()] = f.Name()
	}
	// NOTE: Ignoring potential errors from input.Err()
}
