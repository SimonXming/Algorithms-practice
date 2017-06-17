package main

import (
	"github.com/SimonXming/Algorithms-practice/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSorted(t *testing.T) {
	l := &lib.IntSlice{1, 6, 3}

	// assert equality
	is_sorted := l.IsSorted()
	println(is_sorted)
	assert.Equal(t, is_sorted, true, "they should be equal")

}
