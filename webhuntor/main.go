package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Scrapping the web")

	scraperUrl := "https://wiki.hackerspaces.org/Talk:Communication"

	// initialize the collector
	col := colly.NewCollector(
		// allow only domains
		colly.AllowedDomains("wiki.hackerspaces.org", "www.wiki.hackerspaces.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	//on every html element which has h1 elment
	col.OnHTML("h1.firstHeading", func(h *colly.HTMLElement) {
		// print text
		fmt.Printf("Found H1: %s\n", h.Text)
	})

	err := col.Visit(scraperUrl)
	if err != nil {
		fmt.Println("Error visiting the URL:", err)
	}

} // end of main
