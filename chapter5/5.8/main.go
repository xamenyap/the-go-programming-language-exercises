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

	elem := ElementByID(doc, "some-id")
	if elem != nil {
		fmt.Println(elem.Data)
		fmt.Println(elem.Attr)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node
	elementByID := func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				found = n
				return false
			}
		}

		return true
	}

	forEachNode(doc, elementByID, elementByID)

	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if cont := pre(n); !cont {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if cont := post(n); !cont {
			return
		}
	}
}
