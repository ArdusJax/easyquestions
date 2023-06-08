package algoexpert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSubsequenceSucceeds(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	actual := IsValidSubsequence(array, sequence)
	assert.True(t, actual, "Subsequence should be found in the slice")
}

func TestIsValidSubsequenceFails(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1000, 677, -187, 0}

	actual := IsValidSubsequence(array, sequence)
	assert.False(t, actual, "Subsequence should not be found in the slice")
}
