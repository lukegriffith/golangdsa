package roundrobinIterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
)

func TestNewRRIterator(t *testing.T) {
	_ = New(input)
}

func TestNextRRIterator(t *testing.T) {
	rr := New(input)
	results := []int{}
	for {
		el := rr.Next()
		if el == nil {
			break
		}
		results = append(results, el.Val)
	}
	assert.Equal(t, []int{1, 4, 7, 2, 5, 8, 3, 6, 9}, results)
}
