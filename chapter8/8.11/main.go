package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var done = make(chan struct{})

func main() {
	r := fetchMulti(os.Args[1:])
	close(done)
	if r.err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", r.url, r.err)
		return
	}
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", r.url, r.filename, r.n)
}

type response struct {
	url      string
	filename string
	n        int64
	err      error
}

// fetchMulti fetches name and length of multiple urls,
// but return only the fastest one.
func fetchMulti(urls []string) response {
	responses := make(chan response, len(urls))
	for _, url := range urls {
		url := url
		go func() {
			var r response
			r.filename, r.n, r.err = fetch(url)
			responses <- r
		}()
	}

	return <-responses
}

func fetch(url string) (filename string, n int64, err error) {
	ctx, cancelFn := context.WithCancel(context.Background())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		cancelFn()
		return "", 0, err
	}

	select {
	case <-done:
		cancelFn()
		return "", 0, errors.New("request cancelled")
	default:
		defer cancelFn()
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", 0, err
		}
		defer resp.Body.Close()

		local := path.Base(resp.Request.URL.Path)
		if local == "/" {
			local = "index.html"
		}
		f, err := os.Create(local)
		if err != nil {
			return "", 0, err
		}
		n, err = io.Copy(f, resp.Body)
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		return local, n, err
	}
}
