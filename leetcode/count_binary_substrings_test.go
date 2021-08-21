package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBinarySubstrings(t *testing.T) {
	s1 := "00110011"
	n1 := countBinarySubstrings(s1)
	assert.Equal(t, 6, n1)

	s2 := "10101"
	n2 := countBinarySubstrings(s2)
	assert.Equal(t, 4, n2)
}
