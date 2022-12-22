package examples

import (
	"fmt"
	"testing"
)

type Tree struct {
	root   *Node
	height int
}

func (t *Tree) Close() Tree {
	c := *t
	return c
}

type Node struct {
	val int
}

func Test_Copy(t *testing.T) {
	t1 := Tree{
		root: &Node{
			val: 15,
		},
		height: 1,
	}

	t2 := t1.Close()
	fmt.Println(t2.root.val, " ", t2.height)
	t1.root.val = 18
	t1.height = 3

	fmt.Println(t2.root.val, " ", t2.height)
}
