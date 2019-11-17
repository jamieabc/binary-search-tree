package main

import (
	"fmt"
)

type BST interface {
	Search(int) (*node, *node, error)
	Insert(int)
	Remove(int)
	fmt.Stringer
}

var (
	ErrNotExist = fmt.Errorf("item not exist")
)

type node struct {
	left  *node
	right *node
	val   int
}

func (n *node) Search(target int) (*node, *node, error) {
	tmp := n
	parent := n
	for tmp != nil {
		if tmp.val == target {
			return tmp, parent, nil
		} else {
			parent = tmp
			if tmp.val < target {
				tmp = tmp.right
			} else {
				tmp = tmp.left
			}
		}
	}
	return nil, nil, ErrNotExist
}

func (n *node) Insert(target int) {
	if n.val == 0 {
		n.val = target
		return
	}

	tmp := n
	for tmp != nil {
		if tmp.val < target {
			if tmp.right != nil {
				tmp = tmp.right
			} else {
				tmp.right = &node{
					val: target,
				}
				return
			}
		} else {
			if tmp.left != nil {
				tmp = tmp.left
			} else {
				tmp.left = &node{
					val: target,
				}
				return
			}
		}
	}
}

func (n *node) Remove(target int) {
	toRemove, toRemoveParent, err := n.Search(target)
	if nil != err {
		fmt.Printf("%d not exist\n", target)
		return
	}

	left := false
	if toRemoveParent.left == toRemove {
		left = true
	}

	// no child
	if toRemove.left == nil && toRemove.right == nil {
		if left {
			toRemoveParent.left = nil
		} else {
			toRemoveParent.right = nil
		}
		return
	}

	// 1 child
	if toRemove.left != nil && toRemove.right == nil {
		if left {
			toRemoveParent.left = toRemove.left
		} else {
			toRemoveParent.right = toRemove.left
		}
		return
	}

	if toRemove.left == nil && toRemove.right != nil {
		if left {
			toRemoveParent.left = toRemove.right
		} else {
			toRemoveParent.right = toRemove.right
		}
		return
	}

	// 2 children
	// choose minimum toRemove in right (can also choose maximum in left)
	tmp := minimum(n.right)
	_, newNodeParent, _ := n.Search(tmp.val)
	newNodeParent.left = nil
	if left {
		toRemoveParent.left = tmp
	} else {
		toRemoveParent.right = tmp
	}
	tmp.right = toRemove.right
}

func minimum(n *node) *node {
	if n.left == nil && n.right == nil {
		return n
	}

	lMin := minimum(n.left)
	rMin := minimum(n.right)

	if lMin.val < rMin.val {
		return lMin
	} else {
		return rMin
	}
}

func (root *node) String() string {
	return traverse(root)
}

func traverse(n *node) string {
	if n == nil {
		return " nil"
	}

	return fmt.Sprintf("%d %s %s", n.val, traverse(n.left), traverse(n.right))
}

func newBST(data []int) BST {
	n := &node{}

	for _, d := range data {
		n.Insert(d)
	}
	return n
}
