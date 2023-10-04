package main

import (
	"log"
	"os"
	"path"
	"plugin"

	"github.com/mrchi/golang/Black-Hat-Go/ch10/10.03-plugin/scanner"
)

const (
	PLUGINS_DIR = "../plugins/"
	HOST        = "127.0.0.1"
	PORT        = 8000
)

func main() {
	files, err := os.ReadDir(PLUGINS_DIR)
	if err != nil {
		log.Fatalln(err)
	}
	if len(files) == 0 {
		log.Fatalln("No plugins.")
	}

	for _, file := range files {
		log.Printf("Found plugin: %s\n", file.Name())

		p, err := plugin.Open(path.Join(PLUGINS_DIR, file.Name()))
		if err != nil {
			log.Fatalln(err)
		}

		n, err := p.Lookup("New")
		if err != nil {
			log.Fatalln(err)
		}
		newFunc, ok := n.(func() scanner.Checker)
		if !ok {
			log.Fatalln("Plugin entry point is invalid")
		}

		checker := newFunc()
		res := checker.Check(HOST, PORT)
		if res.Vulnerable {
			log.Printf("Host is vulnerable: %s\n", res.Details)
		} else {
			log.Println("Host is not vulnerable")
		}
	}
}
