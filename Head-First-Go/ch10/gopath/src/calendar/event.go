package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 10 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string {
	return e.title
}
