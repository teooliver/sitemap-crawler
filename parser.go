package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

// Parser defines the parsing interface
type Parser interface {
	GetSeoData(resp *http.Response) (SeoData, error)
}

// DefaultParser is en empty struct for implmenting default parser
type DefaultParser struct {
}

// GetSeoData concrete implementation of the default parser
func (d DefaultParser) GetSeoData(res *http.Response) (SeoData, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return SeoData{}, err
	}
	result := SeoData{}
	result.URL = res.Request.URL.String()
	result.StatusCode = res.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]").Attr("content")
	return result, nil
}
