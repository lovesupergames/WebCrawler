package htmlURL

import (
	"fmt"
	"net/url"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	// выходим если имя домена у базы и курент не совпадают
	parseBase, errParse := url.Parse(rawBaseURL)
	if errParse != nil {
		fmt.Println(errParse)
		return
	}
	parseCurrent, errParse := url.Parse(rawCurrentURL)
	if errParse != nil {
		fmt.Println(errParse)
		return
	}
	if parseBase.Host != parseCurrent.Host {
		return
	}

	//нормализовали и добавили курентЮРЛ в паджес
	normUrl, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
	}
	// increment if visited
	if _, visited := pages[normUrl]; visited {
		pages[normUrl]++
		return
	}
	pages[normUrl] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	// получили содержание страницы
	body, errHTML := GetHTML(rawCurrentURL)
	if errHTML != nil {
		fmt.Println(errHTML)
		return
	}
	//распарсили на ссылки
	urlSlice, errURLS := GetURLsFromHTML(body, rawCurrentURL)
	if errURLS != nil {
		fmt.Println(errURLS)
	}

	//	//добавили все ссылки
	//	for _, link := range urlSlice {
	//		parsedLink, errParse := url.Parse(link)
	//		if errParse != nil {
	//			fmt.Println(errParse)
	//			continue
	//		}
	//		if parsedLink.Host == parseBase.Host {
	//			normLink, errLink := NormalizeURL(link)
	//			if errLink != nil {
	//				fmt.Println(errLink)
	//				return
	//			}
	//
	//			if _, ok := pages[normLink]; !ok {
	//				pages[normLink] = 1
	//			} else {
	//				pages[normLink]++
	//			}
	//		}
	//
	//	}
	//
	//	//получили ссылку из мапы
	//	for key, value := range pages {
	//		// выходим если уже были в этой ссылке
	//
	//		if value > 1 {
	//			continue
	//		}
	//
	//		normalURL, errURL := NormalizeURL(key)
	//		if errURL != nil {
	//			fmt.Println(errURL)
	//			return
	//		}
	//		CrawlPage(rawBaseURL, normalURL, pages)
	//
	//	}
	//}
	for _, nextURL := range urlSlice {
		CrawlPage(rawBaseURL, nextURL, pages)
	}
}
