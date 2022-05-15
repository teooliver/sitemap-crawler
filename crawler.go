package main

import (
	"log"
	"net/http"
)

func scrapePage(url string, token chan struct{}, parser Parser) (SeoData, error) {
	res, err := crawlPage(url, token)
	if err != nil {
		return SeoData{}, err
	}
	data, err := parser.GetSeoData(res)
	if err != nil {
		return SeoData{}, err
	}
	return data, nil
}

func crawlPage(url string, tokens chan struct{}) (*http.Response, error) {
	tokens <- struct{}{}
	resp, err := makeRequest(url)
	<-tokens
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ScrapeSitemap scrapes a given sitemap
func scrapeSitemap(url string, parser Parser, concurrency int) []SeoData {
	results := extractSitemapURLs(url)
	res := scrapeUrls(results, parser, concurrency)
	return res
}

func scrapeUrls(urls []string, parser Parser, concurrency int) []SeoData {
	tokens := make(chan struct{}, concurrency)
	var n int
	n++
	worklist := make(chan []string)
	results := []SeoData{}
	go func() { worklist <- urls }()
	for ; n > 0; n-- {
		list := <-worklist
		for _, url := range list {
			if url != "" {
				n++
				go func(url string, token chan struct{}) {
					// log.Printf("Requesting URL: %s", url)
					res, err := scrapePage(url, tokens, parser)
					if err != nil {
						log.Printf("Encountered error, URL: %s", url)
						log.Printf("Error: %s", err)
					} else {
						results = append(results, res)
					}
					worklist <- []string{}
				}(url, tokens)
			}
		}
	}
	return results
}
