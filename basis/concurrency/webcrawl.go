package main

import (
	"fmt"
	"sync"
)

// Fetcher 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type safeState struct {
	v   map[string]bool
	mux sync.Mutex
}

func (c *safeState) setState(key string, state bool) {
	c.mux.Lock()
	c.v[key] = state
	c.mux.Unlock()
}

func (c *safeState) value(key string) (bool, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	v, ok := c.v[key]
	return v, ok
}

var (
	// urlm     = make(map[string]bool)
	urlState = safeState{v: make(map[string]bool)}
)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	defer func() {
		ch <- 1
	}()
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}

	// 并发安全的实现
	if _, ok := urlState.value(url); ok {
		return
	}

	urlState.setState(url, false)
	defer urlState.setState(url, true)
	// 并发安全的实现结束

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	subCh := make(chan int)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, subCh)
	}
	for i := 0; i < len(urls); i++ {
		<-subCh
	}
	return
}

func main() {
	ch := make(chan int)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
