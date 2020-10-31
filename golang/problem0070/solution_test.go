package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	outcome := PerfectNumber(1)
	expected := 19

	assert.Equal(t, expected, outcome)
}

func TestExample2(t *testing.T) {
	outcome := PerfectNumber(2)
	expected := 28
	assert.Equal(t, expected, outcome)
}
