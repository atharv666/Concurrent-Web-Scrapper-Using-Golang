package main

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

func fetchUrl(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("%s -> Error: %v", url, err)
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("%s -> Failed to Parse Html", url)
		return
	}

	var title string

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if title == "" {
		title = "No title Found"
	}

	ch <- fmt.Sprintf("%s -> %s", url, title)
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://www.geeksforgeeks.org",
		"https://www.google.com",
		"https://www.microsoft.com",
		"https://www.firefox.com",
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg, ch) //goroutine for concurrency
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}
}
