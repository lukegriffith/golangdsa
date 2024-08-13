package queue

type Queue[T any] struct {
	content []T
}

func (q Queue[T]) Add(input T) {
	q.content = append(q.content, input)
}

func (q Queue[T]) Pop() *T {

	l := len(q.content)
	if l == 0 {
		return nil
	}
	i := q.content[len(q.content)-1]
	if len(q.content) == 1 {
		q.content = []T{}
	} else {
		q.content = q.content[0 : len(q.content)-2]
	}
	return &i
}

func New[T any]() Queue[T] {
	return Queue[T]{content: []T{}}
}
