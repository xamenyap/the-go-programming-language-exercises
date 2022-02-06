package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Enter search term as the second argument")
		return
	}

	var xkcds []xkcd
	nums := []int{570, 571, 572}
	const url = "https://xkcd.com/%d/info.0.json"

	for _, n := range nums {
		resp, err := http.Get(fmt.Sprintf(url, n))
		if err != nil {
			log.Fatal(err)
		}

		var xkcd xkcd
		if err := json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()
		xkcds = append(xkcds, xkcd)
	}

	search := os.Args[1]
	found := make([]xkcd, 0)
	for _, xkcd := range xkcds {
		if strings.Contains(xkcd.Title, search) || strings.Contains(xkcd.Transcript, search) {
			found = append(found, xkcd)
		}
	}

	fmt.Printf("Num\tYear\tTitle\tImg\n")
	for _, xkcd := range found {
		fmt.Printf("%d\t%s\t%s\t%s\n", xkcd.Num, xkcd.Year, xkcd.Title, xkcd.Img)
	}
}

type xkcd struct {
	Month      string `json:"month"`
	Day        string `json:"day"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}
