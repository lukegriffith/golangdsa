package queue

type Queue[T any] struct {
	content []T
}

func (q *Queue[T]) Add(input T) {
	q.content = append(q.content, input)
}

func (q *Queue[T]) Pop() T {

	l := len(q.content)
	i := q.content[l-1]
	if len(q.content) == 1 {
		q.content = []T{}
	} else {
		q.content = q.content[0 : l-2]
	}
	return i
}

func New[T any]() *Queue[T] {
	return &Queue[T]{content: []T{}}
}
