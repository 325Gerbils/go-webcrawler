package main

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/jackdanger/collectlinks"
)

// Crawler struct
type Crawler struct {
	visited map[string]bool
	queue   chan string
}

// CrawlFunc - Crawls starting from a URL and executes function run on every found url
func (c Crawler) CrawlFunc(initialURL string, run func(url string)) {

	c.queue = make(chan string)
	c.visited = make(map[string]bool)

	go func() {
		c.queue <- initialURL
	}()
	for uri := range c.queue {
		c.enqueue(uri, c.queue)
		go run(uri)
	}
}

func (c Crawler) enqueue(uri string, queue chan string) {
	c.visited[uri] = true
	links := getURLsFromPage(uri)
	for _, link := range links {
		absolute := fixURL(link, uri)
		if uri != "" && !c.visited[absolute] {
			go func() { c.queue <- absolute }()
		}
	}
}

func getURLsFromPage(url string) []string {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := http.Client{Transport: transport}

	resp, err := client.Get(url)
	if err != nil {
		return []string{""}
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	return links
}

func fixURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	return uri.String()
}

// GetFound found urls
func (c Crawler) GetFound() (out []string) {
	for k, v := range c.visited {
		if v {
			out = append(out, k)
		}
	}
	return
}
