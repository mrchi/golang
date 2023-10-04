package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/mrchi/golang/Black-Hat-Go/ch10/10.03-plugin/scanner"
)

var errKeywords = []string{"syntax"}

type FTPChecker struct {
	errRegexes []*regexp.Regexp
	payloads   []string
}

func (c *FTPChecker) Check(host string, port int) *scanner.Result {
	client := http.Client{}

	for _, payload := range c.payloads {
		form := url.Values{}
		form.Add("uname", payload)
		form.Add("passwd", "p")
		form.Add("submit", "Submit")

		reqUrl := fmt.Sprintf("http://%s:%d/Less-11/", host, port)
		resp, err := client.PostForm(reqUrl, form)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()

		bodyStr := string(body)

		for idx, re := range c.errRegexes {
			if re.MatchString(bodyStr) {
				return &scanner.Result{
					Vulnerable: true,
					Details:    fmt.Sprintf("SQL error found (%#v) for payload: %#v\n", errKeywords[idx], payload),
				}

			}
		}
	}
	return &scanner.Result{Vulnerable: false}
}

func New() scanner.Checker {
	var errRegexes []*regexp.Regexp
	for _, e := range errKeywords {
		re := regexp.MustCompile(fmt.Sprintf(".*%s.*", e))
		errRegexes = append(errRegexes, re)
	}

	payloads := []string{"baseline", ")", "(", "\"", "'"}

	return &FTPChecker{errRegexes: errRegexes, payloads: payloads}
}
