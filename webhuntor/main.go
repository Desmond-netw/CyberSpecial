package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {

	// WEB SCRAPER TITLE
	fmt.Println("Web Scraper")

	url := "https://www.trackingdifferences.com/ETF/ISIN/IE00B1XNHC34"
	// Building web collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"),
		colly.UserAgent("Mozilla/5.0 (windows NT 10.0; Win64; x64)"),
	)
	// print the url that we search
	c.OnRequest(func(rq *colly.Request) {
		fmt.Println("Visiting", rq.URL.String())
	})
	// what we looking for
	c.OnHTML("h1.page-title", func(e *colly.HTMLElement) {
		fmt.Println("Header Titile :", e.Text)
	})
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Printf("Link found : %q -> %s\n", h.Text, link)
	})

	// visit lurl
	err := c.Visit(url)
	if err != nil {
		log.Fatal("Error visiting url:", err)
	}

}
