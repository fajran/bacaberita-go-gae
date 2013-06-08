package parser

import (
	"errors"
	"fmt"
	"time"
)

func ParseDate(str string) (time.Time, error) {
	formats := [...]string{
		"Mon, 02 Jan 2006 15:04:05 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano}
	for _, format := range formats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, nil
		}
	}

	msg := fmt.Sprintf("Unable to parse date: %s", str)
	return time.Time{}, errors.New(msg)
}
