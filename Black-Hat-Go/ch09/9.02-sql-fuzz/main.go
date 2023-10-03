// 靶场可以使用 https://github.com/Audi-1/sqli-labs
// docker run -dt --name sqli-lab -p [PORT]:80 acgpiano/sqli-labs:latest

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const (
	TARGET_URL = "http://localhost:8000/Less-11/"
)

var errRegexes []*regexp.Regexp
var errKeywords = []string{"syntax"}

func init() {

	for _, e := range errKeywords {
		re := regexp.MustCompile(fmt.Sprintf(".*%s.*", e))
		errRegexes = append(errRegexes, re)
	}
}

func main() {
	payloads := []string{"baseline", ")", "(", "\"", "'"}
	client := http.Client{}

	for _, payload := range payloads {
		form := url.Values{}
		form.Add("uname", payload)
		form.Add("passwd", "p")
		form.Add("submit", "Submit")

		resp, err := client.PostForm(TARGET_URL, form)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()

		bodyStr := string(body)

		for idx, re := range errRegexes {
			if re.MatchString(bodyStr) {
				fmt.Printf("[*]SQL error found (%#v) for payload: %#v\n", errKeywords[idx], payload)
			}
		}
	}
}
