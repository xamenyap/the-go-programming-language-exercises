package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	appearances := make(map[string]fileNames)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, appearances)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, appearances)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(appearances[line].files(), " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearances map[string]fileNames) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		counts[txt]++
		if _, ok := appearances[txt]; !ok {
			appearances[txt] = make(fileNames)
		}
		appearances[txt][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

type fileNames map[string]bool

func (fn fileNames) files() []string {
	files := make([]string, 0)
	for f := range fn {
		files = append(files, f)
	}

	return files
}
