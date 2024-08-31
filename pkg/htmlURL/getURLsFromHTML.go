package htmlURL

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	if htmlBody == "" {
		return []string{}, errors.New("empty html body")
	}
	if rawBaseURL == "" {
		return []string{}, errors.New("empty raw base URL")
	}

	linkSlice := []string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					url := a.Val
					if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
						url = rawBaseURL + url
					}
					linkSlice = append(linkSlice, url)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Создаем парсер для HTML
	body, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	// Обходим дерево и вызываем рекурсивную функцию
	f(body)

	return linkSlice, nil
}
