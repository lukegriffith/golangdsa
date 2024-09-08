package tree

type Tree[T any] struct {
	content T
	left    *Tree[T]
	right   *Tree[T]
}
