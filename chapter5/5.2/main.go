package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse html: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	visit(m, doc)

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func visit(m map[string]int, n *html.Node) {
	if n != nil {
		if n.Type == html.ElementNode {
			m[n.Data]++
		}

		visit(m, n.NextSibling)
		visit(m, n.FirstChild)
	}
}
