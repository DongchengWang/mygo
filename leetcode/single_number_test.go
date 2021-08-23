package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	got1 := singleNumber([]int{2, 2, 1})
	assert.Equal(t, 1, got1)

	got2 := singleNumber([]int{4, 1, 2, 1, 2})
	assert.Equal(t, 4, got2)

	got3 := singleNumber([]int{1})
	assert.Equal(t, 1, got3)
}
