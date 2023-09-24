package dbminer

import (
	"fmt"
	"regexp"
)

type Table struct {
	Name    string
	Columns []string
}

type Database struct {
	Name   string
	Tables []Table
}

type Schema struct {
	Databases []Database
}

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)pass(word)?`),
		regexp.MustCompile(`(?i)hash`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)card`),
		regexp.MustCompile(`(?i)security`),
		regexp.MustCompile(`(?i)key`),
	}
}

func Search(m DatabaseMiner) error {
	s, err := m.GetSchema()
	if err != nil {
		return err
	}

	re := getRegex()
	for _, database := range s.Databases {
		for _, table := range database.Tables {
			for _, field := range table.Columns {
				for _, r := range re {
					if r.MatchString(field) {
						fmt.Println(database.Name, table.Name)
						fmt.Printf("[+] HIT: %s\n", field)
					}
				}
			}
		}
	}

	return nil
}
