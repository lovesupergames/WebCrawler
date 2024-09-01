package main

import (
	"fmt"
	"github.com/lovesupergames/WebCrawler/pkg/htmlURL"
	_ "github.com/lovesupergames/WebCrawler/pkg/htmlURL"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	maxConcurrency, errMC := strconv.Atoi(os.Args[2])
	if errMC != nil && maxConcurrency < 0 {
		fmt.Println("Wrong concurrency provided")
		return
	}
	maxPages, errMP := strconv.Atoi(os.Args[3])
	if errMP != nil && maxPages <= 0 {
		fmt.Println("Wrong maxPages provided")
		return
	}

	c := htmlURL.Config{
		MaxPages:           maxPages,
		Pages:              make(map[string]int),
		BaseURL:            baseUrl,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, maxConcurrency), // Example value for max concurrency
		Wg:                 &sync.WaitGroup{},
	}

	c.Wg.Add(1)
	go c.CrawlPage(rawBaseURL)
	c.Wg.Wait()

	htmlURL.PrintReport(c.Pages, baseUrl.String())
}
