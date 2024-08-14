package queue

type Queue[T any] struct {
	content []T
}

func (q *Queue[T]) Add(input T) {
	q.content = append(q.content, input)
}

func (q *Queue[T]) Pop(idx int) T {

	l := len(q.content)
	i := q.content[idx]
	if len(q.content) <= 1 {
		q.content = []T{}
	} else {
		q.content = append(q.content[0:idx], q.content[idx+1:l]...)
	}
	return i
}

func (q *Queue[T]) Length() int {
	return len(q.content)
}

func New[T any]() *Queue[T] {
	return &Queue[T]{content: []T{}}
}
