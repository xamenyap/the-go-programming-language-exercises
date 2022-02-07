package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	search := flag.String("s", "", "movie to search for poster")
	flag.Parse()

	if *search == "" {
		flag.PrintDefaults()
		return
	}

	posterURL := queryPoster(*search)
	if posterURL != "" {
		downloadPoster(posterURL)
	}
}

func queryPoster(search string) string {
	f, err := os.Open("./api_key.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	apiKey, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "http://www.omdbapi.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("apikey", string(apiKey))
	q.Add("t", search)
	q.Add("type", "movie")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var r result
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatal(err)
	}

	return r.Poster
}

func downloadPoster(posterURL string) {
	fileName := strings.Replace(path.Base(posterURL), " ", "_", -1)
	resp, err := http.Get(posterURL)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.OpenFile(fileName+".tmp", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		out.Close()
		log.Fatal(err)
	}

	w := new(progressWriter)
	io.Copy(out, io.TeeReader(resp.Body, w))
	out.Close()

	// move to new line after printing progress
	fmt.Println()

	if err := os.Rename(fileName+".tmp", fileName); err != nil {
		log.Fatal(err)
	}
}

type result struct {
	Title  string `json:"title"`
	Poster string `json:"poster"`
}

type progressWriter struct {
	total uint64
}

func (w *progressWriter) Write(p []byte) (n int, err error) {
	w.total += uint64(len(p))
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloaded %d bytes", w.total)
	return len(p), nil
}
