package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	moreThanOneYear := make([]*github.Issue, 0)
	lessThanOneYear := make([]*github.Issue, 0)
	lessThanOneMonth := make([]*github.Issue, 0)
	now := time.Now()

	for _, item := range result.Items {
		if now.Sub(item.CreatedAt) >= 365*24*time.Hour {
			moreThanOneYear = append(moreThanOneYear, item)
		} else if now.Sub(item.CreatedAt) >= 30*24*time.Hour {
			lessThanOneYear = append(lessThanOneYear, item)
		} else {
			lessThanOneMonth = append(lessThanOneMonth, item)
		}
	}

	fmt.Println("\nMore Than One Year Old")
	for _, item := range moreThanOneYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
	}

	fmt.Println("\nLess Than One Year Old")
	for _, item := range lessThanOneYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
	}

	fmt.Println("\nLess Than One Month Old")
	for _, item := range lessThanOneMonth {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02"))
	}
}
