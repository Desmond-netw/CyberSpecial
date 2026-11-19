package main

import (
	"fmt"
	"net/http"
	"time"
)

type Platform struct {
	Name    string
	URL     string // profile URL template with %s for username
	Method  string // "api" or "profile"
	Exists  bool
	Profile string
}

var platforms = []Platform{
	{Name: "GitHub", URL: "https://api.github.com/users/%s", Method: "api"},
	{Name: "X", URL: "https://x.com/%s", Method: "profile"},
	{Name: "Instagram", URL: "https://www.instagram.com/%s/", Method: "profile"},
	{Name: "Facebook", URL: "https://www.facebook.com/%s", Method: "profile"},
	{Name: "LinkedIn", URL: "https://www.linkedin.com/in/%s", Method: "profile"},
	{Name: "Github", URL: "https://github.com/%s", Method: "profile"},
}

func main() {

	// load env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	username := "desmond-netw" // target user

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	client.Transport = &http.Transport{
		Proxy: nil, // no proxy needed
	}

	fmt.Printf("Checking username %q\n\n", username)

	for _, p := range platforms {
		url := fmt.Sprintf(p.URL, username)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; UsernameChecker/1.0)")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%-10s Error %v\n", p.Name, err)
			continue
		}

		// Determine if account exist base on status code
		// 200 == exist, 404 = not found, 302/303 = redirect ( often to login )

		exist := resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusFound ||
			resp.StatusCode == http.StatusSeeOther

		status := "Not Found"
		if exist {
			status = "FOUND"
		}
		fmt.Printf("%-10s %s (HTTP %d)\n", p.Name, status, resp.StatusCode)
	}

}
