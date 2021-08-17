package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	got := reverseVowels("leetcode")
	want := "leotcede"
	assert.Equal(t, want, got)
}
