package main

import (
	"encoding/json"
	"fmt"
	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var parser *argparse.Parser = argparse.NewParser("goad", "an HTTP load testing tool")

func main() {
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatalf("couldn't parse arguments %v", err)
	}
	testCooldownSeconds := time.Duration(*testCooldown) * time.Second
	suite := TestSuite{
		Start:    time.Now(),
		Steps:    *testSteps,
		Cooldown: testCooldownSeconds,
		Repeats:  *testRepeats,
		URL:      *requestUrlString,
		Args:     os.Args,
		Tests:    []TestRepeat{},
	}
	suite.Commit, err = ParseGit()
	if err != nil {
		log.Fatalf("could no parse Git: %v", err)
	}
	suite.RunTests()
	suiteJSON, err := json.Marshal(suite)
	fmt.Printf("%+s", suiteJSON)
}
