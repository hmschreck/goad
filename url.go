package main

import (
	"errors"
	"github.com/akamensky/argparse"
	"net/url"
)

var requestUrlString *string = parser.String("", "url", &argparse.Options{
	Required: true,
	Help:     "URL to hit for testing",
	Validate: ValidateURL,
})

func ValidateURL(args []string) (err error) {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("too many args for request URL")
	}
	_, err = url.Parse(args[0])
	return
}
