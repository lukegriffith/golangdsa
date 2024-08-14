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
	q.Add(2)
	i := q.Pop(0)
	assert.Equal(t, 1, i)
	i = q.Pop(0)
	assert.Equal(t, 2, i)
	assert.Equal(t, []int{}, q.content)
}

func TestQueueLength(t *testing.T) {
	q := New[int]()
	q.Add(1)
	assert.Equal(t, 1, q.Length())
	q.Add(1)
	assert.Equal(t, 2, q.Length())
}
