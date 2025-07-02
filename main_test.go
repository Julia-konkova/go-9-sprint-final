package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {

	tests := []struct {
		size   int
		expect int
	}{
		{10, 10},
		{0, 0},
		{-5, 0},
		{1_000_000, 1_000_000},
	}

	for _, test := range tests {
		actual := generateRandomElements(test.size)
		require.Len(t, actual, test.expect)
	}
}

func TestMaximum(t *testing.T) {

	tests := []struct {
		slice     []int
		expectMax int
	}{
		{nil, 0},
		{[]int{10}, 10},
		{[]int{-10, -5}, -5},
		{[]int{991, 67, 36, 241, 651, 108, 634, 591, 2, 819}, 991},
	}
	for _, test := range tests {
		max := maximum(test.slice)
		assert.Equal(t, test.expectMax, max)
	}
}
