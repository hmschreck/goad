package main

import "github.com/akamensky/argparse"

var requestHeaders *[]string = parser.StringList("H", "header", &argparse.Options{
	Required: false,
	Help:     "headers to attach to the HTTP request",
	Default:  nil,
})

