package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p := DefaultParser{}
	results := scrapeSitemap("https://www.hollandandbarrett.com/sitemap-healthhub.xml", p, 10)
	for _, res := range results {
		fmt.Println(res)
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
