package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	outcome := Solution([]int{1, 2, 3, 4, 5})
	expected := []int{120, 60, 40, 30, 24}

	assert.Equal(t, expected, outcome)
}

func TestExample2(t *testing.T) {
	outcome := Solution([]int{3, 2, 1})
	expected := []int{2, 3, 6}

	assert.Equal(t, expected, outcome)
}
