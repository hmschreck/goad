package main

import (
	"github.com/akamensky/argparse"
	"io/ioutil"
	"os"
	"strings"
)

var requestPayload []byte
var requestData *string = parser.String("d", "data", &argparse.Options{
	Required: false,
	Help:     "HTTP data to include with request",
	Default:  nil,
})

func ParseBody() (err error) {
	if strings.HasPrefix(*requestData, "@") {
		data := strings.TrimPrefix(*requestData, "@")
		loadFile, err := os.Open(data)
		if err != nil {
			return err
		}
		requestPayload, err = ioutil.ReadAll(loadFile)
		if err != nil {
			return err
		}
	} else {
		requestPayload = []byte(*requestData)
	}
	return
}
