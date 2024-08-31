package main

import (
	"fmt"
	"github.com/lovesupergames/WebCrawler/pkg/htmlURL"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	pages := make(map[string]int)

	htmlURL.CrawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
