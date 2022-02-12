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
	visit(doc)
}

func visit(n *html.Node) {
	if n != nil {
		if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
			fmt.Println(n.Data)
		}

		visit(n.NextSibling)
		visit(n.FirstChild)
	}
}
