package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	words := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[strings.TrimSpace(sc.Text())]++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words\tcount\n")
	for s, n := range words {
		fmt.Printf("%s\t%d\n", s, n)
	}
}
