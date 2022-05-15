package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Starting script %s\n", start)
	p := DefaultParser{}
	results := scrapeSitemap("https://www.omgubuntu.co.uk/sitemap-1.xml", p, 10)

	for i, res := range results {
		fmt.Println(i, res.URL)
	}

	timeElapsed := time.Since(start)
	fmt.Printf("URL amount %d\n", len(results))
	fmt.Printf("Time Elapsed %s\n", timeElapsed)
}

/*
# TODOs:
- Option to recieve the url via cli
- Better display of SeoData results on the terminal
- Generate reports with error codes and urls
*/
