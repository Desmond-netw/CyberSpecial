package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Hunting the web")

	targetUrl := "https://www.trackingdifferences.com/ETF/ISIN/IE00B1XNHC34"

	// initialize the colly collector
	col := colly.NewCollector(
		// allow only domains
		colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	//on every html element which has h1 elment
	col.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// print text
		fmt.Printf("link found: %q => %s \n", e.Text, link)

		// visit link found in page
		col.Visit(e.Request.AbsoluteURL(link))

	})

	err := col.Visit(targetUrl)
	if err != nil {
		fmt.Println("Error visiting the URL:", err)
	}

} // end of main
