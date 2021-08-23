package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	got1 := hammingDistance(1, 4)
	assert.Equal(t, 2, got1)

	got2 := hammingDistance(3, 1)
	assert.Equal(t, 1, got2)
}
