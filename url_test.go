package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateURL(t *testing.T) {
	args := []string{"https://google.com"}
	err := ValidateURL(args)
	assert.Nil(t, err)

	args = []string{}
	err = ValidateURL(args)
	assert.NotNil(t, err)

	args = []string{"https://google.com", "https://sydneythe.dev"}
	err = ValidateURL(args)
	assert.NotNil(t, err)
}
