package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueAdd(t *testing.T) {
	q := New[int]()
	q.Add(1)
}

func TestQueuePop(t *testing.T) {
	q := New[int]()
	q.Add(1)
	i := q.Pop()
	assert.Equal(t, 1, i)
}
