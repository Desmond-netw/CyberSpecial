package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Etf structs
type EtfInfo struct {
	Title              string
	Replication        string
	Earnings           string
	TotalExpenseRatio  string
	TrackingDifference string
	FundSize           string
}

func main() {
	etfinfo := EtfInfo{}

	// WEB SCRAPER TITLE
	fmt.Println("Web Scraper")

	url := "http://www.trackingdifferences.com/ETF/ISIN/IE00B1XNHC34"
	// Building web collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"),
		colly.UserAgent("Mozilla/5.0 (windows NT 10.0; Win64; x64)"),
	)

	// print the url that we search
	c.OnRequest(func(rq *colly.Request) {
		rq.Headers.Set("Accept-Language", "en-us;q=0.9")
		fmt.Println("Visiting", rq.URL.String())
	})

	// what we looking for
	c.OnHTML("h1.page-title", func(e *colly.HTMLElement) {
		fmt.Println("Header Titile :", e.Text)
	})
	// geting more details from page
	c.OnHTML("div.descfloat p.desc", func(h *colly.HTMLElement) {
		// set selection
		selection := h.DOM
		childNode := selection.Children().Nodes
		if len(childNode) == 3 {
			dscrp := cleanDesc(selection.Find("span.desctitle").Text())
			value := selection.FindNodes(childNode[2]).Text()

			switch dscrp {
			case "Replication":
				etfinfo.Replication = value
			case "Earnings":
				etfinfo.Earnings = value
			case "TER":
				etfinfo.TotalExpenseRatio = value
			case "TD":
				etfinfo.TrackingDifference = value
			case "Fund size":
				etfinfo.FundSize = value
				break
			}

		}

	})

	c.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		enc.Encode(etfinfo)
	})

	// error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL", r.Request.URL, "faild with response", r, "\nError:", err)
	})

	// visit lurl
	err := c.Visit(url)
	if err != nil {
		log.Fatal("Error visiting url:", err)
	}

}

func cleanDesc(s string) string {
	return strings.TrimSpace(s)
}
