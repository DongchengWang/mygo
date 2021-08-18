package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	n1 := countSubstrings("aaa")
	assert.Equal(t, 6, n1)

	n2 := countSubstrings("abc")
	assert.Equal(t, 3, n2)
}
