package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	s1 := "egg"
	s2 := "add"
	got := isIsomorphic(s1, s2)
	assert.True(t, got)
}

func TestIsIsomorphic2(t *testing.T) {
	s1 := "foo"
	s2 := "bar"
	got := isIsomorphic(s1, s2)
	assert.False(t, got)
}

func TestIsIsomorphic3(t *testing.T) {
	s1 := "paper"
	s2 := "title"
	got := isIsomorphic(s1, s2)
	assert.True(t, got)
}

func TestIsIsomorphic4(t *testing.T) {
	s1 := "abb"
	s2 := "baa"
	got := isIsomorphic(s1, s2)
	assert.True(t, got)
}
