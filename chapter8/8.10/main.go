package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

type url struct {
	name  string
	depth int
}

var tokens = make(chan struct{}, 20)

const maxDepth = 3

var c = http.Client{
	Timeout: 2 * time.Second,
}

func main() {
	worklist := make(chan url)
	//var wg sync.WaitGroup
	//wg.Add(len(os.Args[1:]))

	var n int
	n = n + len(os.Args[1:])
	go func() {
		for _, arg := range os.Args[1:] {
			worklist <- url{name: arg, depth: 0}
		}
	}()

	seen := make(map[string]bool)
	ctx, cancelFn := context.WithCancel(context.Background())
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancelFn()
	}()

	for ; n > 0; n-- {
		item := <-worklist

		if !seen[item.name] {
			seen[item.name] = true

			newURLs := crawl(ctx, item.name, item.depth+1)
			for _, u := range newURLs {
				n++
				go func(u string, d int) {
					worklist <- url{
						name:  u,
						depth: d,
					}
				}(u, item.depth+1)
			}
		}
	}

	cancelFn()
}

func extract(ctx context.Context, url string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http request from url %s: %w", url, err)
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(ctx context.Context, url string, depth int) []string {
	select {
	case <-ctx.Done():
		return nil
	default:
		if depth > maxDepth {
			return nil
		}

		fmt.Println(url)
		tokens <- struct{}{} // acquire a token
		list, err := extract(ctx, url)
		<-tokens // release the token

		if err != nil {
			log.Print(err)
		}
		return list
	}
}
