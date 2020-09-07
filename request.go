package main

import (
	"errors"
	"github.com/akamensky/argparse"
	"github.com/hmschreck/in"
	"net/http"
)

var requestType *string = parser.String("X", "request", &argparse.Options{
	Required: false,
	Help:     "HTTP method to use",
	Validate: VerifyRequest,
	Default:  http.MethodGet,
})

var methods = []string{http.MethodGet,
	http.MethodDelete,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodOptions,
	http.MethodConnect,
	http.MethodHead,
}

func VerifyRequest(args []string) (err error) {
	if len(args) > 1 {
		return errors.New("too many args for request type")
	}
	isValid, err := in.CheckOne(methods, args[0])
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("value not in list of allowed methods")
	}
	return
}
