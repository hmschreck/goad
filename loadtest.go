package main

import (
	"bytes"
	"net/http"
	"strings"
	"time"
)

func MakeRequest(readyChan chan bool, startChan chan bool, doneChan chan TestResult) (err error) {
	result := TestResult{}
	defer func() {
		doneChan <- result
	}()
	client := http.Client{}
	body := bytes.NewReader(requestPayload)
	req, err := http.NewRequest(*requestType, *requestUrlString, body)
	if err != nil {
		err = err
		return
	}
	for _, header := range *requestHeaders {
		headerSplit := strings.Split(header, ":")
		req.Header.Add(strings.TrimSpace(headerSplit[0]),
			strings.TrimSpace(headerSplit[1]))
	}
	readyChan <- true
	// wait for the synchronized start message
	_ = <-startChan
	result.Start = time.Now()
	resp, err := client.Do(req)
	result.End = time.Now()
	result.Duration = result.End.Sub(result.Start)
	result.Response = resp.StatusCode
	return
}
