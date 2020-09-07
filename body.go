package main

import "github.com/akamensky/argparse"

var requestData *string = parser.String("d", "data", &argparse.Options{
	Required: false,
	Help:     "HTTP data to include with request",
	Default:  nil,
})
