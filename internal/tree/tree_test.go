package tree

import "testing"

func TestNewTree(t *testing.T) {
	_ = Tree[int]{}
}

func TestTree(t *testing.T) {
	tree := Tree[int]{
		1,
		&Tree[int]{
			2,
			nil,
			nil,
		},
		&Tree[int]{
			3,
			nil,
			nil,
		},
	}
	if tree.content != 1 {
		t.Fatal("the content is not right", tree.content)
	}

	if tree.left.content != 2 {
		t.Fatal("left content is not right", tree.left.content)
	}

	if tree.right.content != 3 {
		t.Fatal("right content is not right", tree.right.content)
	}
}
