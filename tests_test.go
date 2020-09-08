package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateSteps(t *testing.T) {
	// test all numbers
	args := []string{"1", "2", "4", "8", "16"}
	err := ValidateSteps(args)
	assert.Nil(t, err)

	// test all letters
	args = []string{"a", "b", "c", "d", "e"}
	err = ValidateSteps(args)
	assert.NotNil(t, err)

	// test mixed letters and numbers
	args = []string{"1", "2", "c", "d", "16"}
	err = ValidateSteps(args)
	assert.NotNil(t, err)

	// for now this should safely return no error.
	args = []string{}
	err = ValidateSteps(args)
	assert.Nil(t, err)
}

func TestValidateRepeats(t *testing.T) {
	args := []string{"1"}
	err := ValidateRepeats(args)
	assert.Nil(t, err)

	args = []string{"0"}
	err = ValidateRepeats(args)
	assert.NotNil(t, err)

	args = []string{"a"}
	err = ValidateRepeats(args)
	assert.NotNil(t, err)

	args = []string{"1", "2"}
	err = ValidateRepeats(args)
	assert.NotNil(t, err)

	args = []string{"-1"}
	err = ValidateRepeats(args)
	assert.NotNil(t, err)
}

func TestValidateCooldown(t *testing.T) {
	args := []string{"3.3"}
	err := ValidateCooldown(args)
	assert.Nil(t, err)

	args = []string{}
	err = ValidateCooldown(args)
	assert.NotNil(t, err)

	args = []string{"-0.1"}
	err = ValidateCooldown(args)
	assert.NotNil(t, err)

	args = []string{"0"}
	err = ValidateCooldown(args)
	assert.NotNil(t, err)

	args = []string{"1", "2"}
	err = ValidateCooldown(args)
	assert.NotNil(t, err)

	args = []string{"a"}
	err = ValidateCooldown(args)
	assert.NotNil(t, err)
}
