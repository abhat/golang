package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SyncUrlMap struct {
	m            sync.Mutex
	visited_urls map[string]bool
}

var visited_urls_map SyncUrlMap

func IsUrlAlreadyVisited(url string) bool {
	visited_urls_map.m.Lock()
	defer visited_urls_map.m.Unlock()
	if _, ok := visited_urls_map.visited_urls[url]; !ok {
		return false
	} else {
		return true
	}
}

func CacheVisitedUrl(url string) {
	visited_urls_map.m.Lock()
	defer visited_urls_map.m.Unlock()
	visited_urls_map.visited_urls[url] = true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		ch <- 1
		return
	}
	CacheVisitedUrl(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	channels := make([]chan int, 0)
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if IsUrlAlreadyVisited(u) == false {
			urlch := make(chan int)
			channels = append(channels, urlch)
			go Crawl(u, depth-1, fetcher, urlch)
		}
	}
	for _, retch := range channels {
		<-retch
	}
	ch <- 1
	return
}

func main() {
	visited_urls_map = SyncUrlMap{visited_urls: make(map[string]bool)}
	ch := make(chan int)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	select {
	case <-ch:
		fmt.Println("Done crawling")
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
