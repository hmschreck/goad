package main

import (
	"errors"
	"github.com/akamensky/argparse"
	"strconv"
)

var testSteps *[]int = parser.IntList("", "steps", &argparse.Options{
	Required: false,
	Help:     "steps for concurrent tests",
	Validate: ValidateSteps,
	Default:  []int{1, 2, 4, 8, 16},
})

func ValidateSteps(args []string) (err error) {
	for _, value := range args {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		if intValue < 1 {
			return errors.New("cannot have a step less than 1")
		}
	}
	return
}

var testCooldown = parser.Float("", "cooldown", &argparse.Options{
	Required: false,
	Help:     "cooldown between test runs",
	Validate: ValidateCooldown,
	Default:  3.0,
})

func ValidateCooldown(args []string) (err error) {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("incorrect arguments for cooldown")
	}
	if cooldown, err := strconv.ParseFloat(args[0], 64); cooldown <= 0 || err != nil {
		return errors.New("could not parse cooldown period")
	}
	return nil
}

var testRepeats = parser.Int("", "repeats", &argparse.Options{
	Required: false,
	Help:     "number of times to repeat the full suite (all steps)",
	Validate: ValidateRepeats,
	Default:  1,
})

func ValidateRepeats(args []string) (err error) {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("too many arguments for repeats")
	}
	if repeats, err := strconv.Atoi(args[0]); err != nil || repeats < 1 {
		return errors.New("could not parse repeats")
	}
	return nil
}
