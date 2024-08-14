package roundrobinIterator

import (
	"github.com/lukegriffith/golangdsa/internal/queue"
)

type RoundRobinIterator[T any] struct {
	content *queue.Queue[*queue.Queue[T]]
}

func New[T any](input [][]T) *RoundRobinIterator[T] {
	q := queue.New[*queue.Queue[T]]()
	for _, inputList := range input {
		if len(inputList) > 0 {
			innerQ := queue.New[T]()
			for _, element := range inputList {
				innerQ.Add(element)
			}
			q.Add(innerQ)
		}
	}
	return &RoundRobinIterator[T]{
		content: q,
	}
}

type IteratorResult[T any] struct {
	Val T
}

func (rr *RoundRobinIterator[T]) Next() *IteratorResult[T] {
	if rr.content.Length() == 0 {
		return nil
	}
	q := rr.content.Pop(0)
	el := q.Pop(0)
	if q.Length() > 0 {
		rr.content.Add(q)
	}
	return &IteratorResult[T]{Val: el}
}
