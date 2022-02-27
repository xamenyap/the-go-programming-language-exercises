package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse html: %v\n", err)
		os.Exit(1)
	}

	words, images := countWordsAndImages(doc)
	fmt.Println("There are", words, "words")
	fmt.Println("There are", images, "images")
}

func countWordsAndImages(n *html.Node) (words, images int) {
	visit(&words, &images, n)
	return
}

func visit(words *int, images *int, n *html.Node) {
	if n != nil {
		if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
			sc := bufio.NewScanner(strings.NewReader(n.Data))
			sc.Split(bufio.ScanWords)
			for sc.Scan() {
				*words++
			}

			if err := sc.Err(); err != nil {
				log.Fatal(err)
			}
		}

		if n.Type == html.ElementNode && n.Data == "img" {
			*images += 1
		}

		visit(words, images, n.NextSibling)
		visit(words, images, n.FirstChild)
	}
}
