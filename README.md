# go-webcrawler
go-webcrawler provides a simple web crawler library for Go

### Installation:

```
go get github.com/325gerbils/go-webcrawler
```

### Usage:

```go
package main

import (
	"fmt"

	crawler "github.com/325gerbils/go-webcrawler"
)

func main() {

	url := "https://github.com"

	c := crawler.Crawler{}
	c.CrawlFunc(url, func(s string) {
		fmt.Println("Found:", s)
	})

}

```

Only crawl URLs matching a given regex:

```go
crawl.FindRegex(url, "(www|http:|https:)+[^\s]+[\w]", func(s string) {})
``` 

Can use `c.GetFound()` to get the list of found URLs as a []string

### Todo:

* Filter URLs to skip image urls, etc
* 
