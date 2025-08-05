
// List duplicates
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	count := make(map[string]int)

	if len(os.Args) < 2 {
		countLines(os.Stdin, count)
	}
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Open: %v\n", err)
			continue
		}
		countLines(file, count)

	}

	for k, v := range count {
		if v > 1 {
			fmt.Printf("%d\t%v\n", v, k)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		count[s.Text()]++
	}
}
