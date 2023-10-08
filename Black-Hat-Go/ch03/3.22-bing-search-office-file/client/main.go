package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/mrchi/golang/Black-Hat-Go/ch03/3.22-bing-search-office-file/metadata"
)

func handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}
	log.Printf("%d: %s\n", i, url)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	coreProps, appProps, err := metadata.NewProperties(r)
	if err != nil {
		return
	}
	log.Printf(
		"%25s %25s - %s %s\n",
		coreProps.Creator,
		coreProps.LastModifiedBy,
		appProps.Application,
		appProps.GetMajorVersion(),
	)
}

func main() {
	domain := os.Args[1]
	fileType := os.Args[2]

	q := fmt.Sprintf("site:%s filetype:%s", domain, fileType)
	search := fmt.Sprintf("https://cn.bing.com/search?q=%s", url.QueryEscape(q))
	log.Printf("Search URL: %s", search)

	req, err := http.NewRequest("GET", search, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc.Find("li.b_algo div.b_title h2").Each(handler)
}
