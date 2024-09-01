package htmlURL

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	MaxPages           int
	Pages              map[string]int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func (cfg *Config) CrawlPage(rawCurrentURL string) {

	cfg.ConcurrencyControl <- struct{}{}

	defer func() {
		<-cfg.ConcurrencyControl
		cfg.Wg.Done()
	}()

	if cfg.PagesLen() >= cfg.MaxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.BaseURL.Hostname() {
		return
	}

	normalizedURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	isFirst := cfg.AddPageVisit(normalizedURL)
	if !isFirst {
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := GetURLsFromHTML(htmlBody, cfg.BaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
	}

	for _, nextURL := range nextURLs {
		cfg.Wg.Add(1)
		go cfg.CrawlPage(nextURL)
	}

}

func (cfg *Config) AddPageVisit(normalizedURL string) (isFirst bool) {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()
	// increment if visited
	if _, visited := cfg.Pages[normalizedURL]; visited {
		cfg.Pages[normalizedURL]++
		return false
	}

	// mark as visited
	cfg.Pages[normalizedURL] = 1
	return true
}

func (cfg *Config) PagesLen() int {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()
	return len(cfg.Pages)
}
