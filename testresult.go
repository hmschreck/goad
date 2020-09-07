package main

import (
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type TestSuite struct {
	Start time.Time `json:"start_time"`
	End time.Time `json:"end_time"`
	Duration time.Duration `json:"duration"`
	Steps []int `json:"steps"`
	Cooldown time.Duration `json:"cooldown"`
	Repeats int `json:"repeats"`
	URL string `json:"url"`
	Args []string `json:"args"`
	Tests []TestRepeat `json:"tests"`
}

func (ts *TestSuite) RunTests() {
	for runs := 0; runs < ts.Repeats; runs ++ {
		repeat := TestRepeat{}
		repeat.Tests = []Test{}
		for _, step := range ts.Steps {
			test := Test{
				Start: time.Now(),
				TestCount: step,
				Results: []TestResult{},
			}
			test.Run()
			test.End = time.Now()
			test.Duration = test.End.Sub(test.Start)
			repeat.Tests = append(repeat.Tests, test)
			time.Sleep(ts.Cooldown)
		}
		ts.Tests = append(ts.Tests, repeat)
		time.Sleep(ts.Cooldown)
	}
}

type TestRepeat struct {
	Tests []Test `json:"tests"`
}

type Test struct {
	Start time.Time `json:"start_time"`
	End time.Time `json:"end_time"`
	Duration time.Duration `json:"duration"`
	TestCount int `json:"count"`
	Minimum time.Duration `json:"minimum"`
	Maximum time.Duration `json:"maximum"`
	StdDev float64 `json:"std_dev"`
	Average float64 `json:"average"`
	Results []TestResult `json:"results"`
}

func (t *Test) Run() {
	log.Debugf("starting test with %d concurrent requests", t.TestCount)
	startChan := make(chan bool)
	doneChan := make(chan TestResult)
	for test := 0; test < t.TestCount; test++ {
		go MakeRequest(startChan, doneChan)
	}
	for test := 0; test < t.TestCount; test++ {
		startChan <- true
	}
	results := []TestResult{}
	for test := 0; test < t.TestCount; test++ {
		result := <- doneChan
		results = append(results, result)
	}
	log.Debugf("%+v", results)
	t.Results = results
	t.Summarize()
}

func (t *Test) Summarize() {
	t.Maximum = t.GetMaximum()
	t.Minimum = t.GetMinimum()
	t.Average = t.GetAverage()
	t.StdDev = t.GetStdDev()
}

func (t Test) GetMaximum() (maximum time.Duration) {
	if len(t.Results) == 0 {
		return
	}
	maximum = t.Results[0].Duration
	for _, result := range t.Results {
		if result.Duration > maximum {
			maximum = result.Duration
		}
	}
	return
}

func (t Test) GetMinimum() (minimum time.Duration) {
	if len(t.Results) == 0 {
		return
	}
	minimum = t.Results[0].Duration
	for _, result := range t.Results {
		if result.Duration < minimum {
			minimum = result.Duration
		}
	}
	return
}

func (t Test) GetAverage() (average float64) {
	if len(t.Results) == 0 {
		return
	}
	var sum float64 = 0
	for _, result := range t.Results {
		sum += float64(result.Duration)
	}
	average = sum / float64(t.TestCount)
	return
}

func (t Test) GetStdDev() (stddev float64) {
	if len(t.Results) == 0 {
		return
	}
	var average float64 = 0
	if t.Average == 0 {
		average = t.GetAverage()
	} else {
		average = t.Average
	}
	var sumOfDiffs float64 = 0
	for _, result := range t.Results {
		sumOfDiffs += math.Pow(float64(result.Duration) - average, 2)
	}
	variance := sumOfDiffs / float64(len(t.Results))
	stddev = math.Sqrt(variance)
	return
}

type TestResult struct {
	Start    time.Time     `json:"start_time"`
	End      time.Time     `json:"end_time"`
	Duration time.Duration `json:"duration"`
	Response int           `json:"response"`
}