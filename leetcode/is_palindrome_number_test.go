package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeNumber(t *testing.T) {
	b1 := isPalindromeNumber(121)
	assert.True(t, b1)
	b2 := isPalindromeNumber(-121)
	assert.False(t, b2)
	b3 := isPalindromeNumber(10)
	assert.False(t, b3)
	b4 := isPalindromeNumber(123)
	assert.False(t, b4)
}
