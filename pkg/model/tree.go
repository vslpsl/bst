package model

import "sort"

type Tree struct {
	root *Node
}

func (t *Tree) Search(key int) (*Node, error) {
	result := t.search(t.root, key)
	if result == nil {
		return nil, ErrNotFound
	}
	return result, nil
}

func (t *Tree) search(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key == root.Key {
		return root
	}
	if key > root.Key {
		return t.search(root.Right, key)
	}
	return t.search(root.Left, key)
}

func NewTree(arr []int) *Tree {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	sort.Ints(arr)
	tree := &Tree{
		root: makeTree(arr),
	}
	return tree
}

func makeTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &Node{Key: arr[0]}
	}
	m := len(arr) / 2
	node := &Node{Key: arr[m]}
	left := makeTree(arr[:m])
	if left != nil {
		left.Parent = node
		node.Left = left
	}
	right := makeTree(arr[m+1:])
	if right != nil {
		right.Parent = node
		node.Right = right
	}
	return node
}

func (t *Tree) Insert(key int) (*Node, error) {
	if t.root == nil {
		t.root = &Node{Key: key}
		return t.root, nil
	}
	return t.insert(t.root, key)
}

func (t *Tree) insert(parent *Node, key int) (*Node, error) {
	if key == parent.Key {
		return nil, ErrAlreadyExists
	}
	if key > parent.Key {
		if parent.Right == nil {
			right := &Node{Key: key, Parent: parent}
			parent.Right = right
			return right, nil
		}
		return t.insert(parent.Right, key)
	}
	if parent.Left == nil {
		left := &Node{Key: key, Parent: parent}
		parent.Left = left
		return left, nil
	}
	return t.insert(parent.Left, key)
}

func (t *Tree) Delete(key int) error {
	if t.search(t.root, key) == nil {
		return ErrNotFound
	}
	t.root = t.delete(t.root, key)
	return nil
}

func (t *Tree) delete(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	if key < root.Key {
		root.Left = t.delete(root.Left, key)
		return root
	}
	if key > root.Key {
		root.Right = t.delete(root.Right, key)
		return root
	}
	if root.Left != nil && root.Right != nil {
		min := t.min(root.Right)
		root.Key = min.Key
		root.Right = t.delete(root.Right, min.Key)
		return root
	}
	if root.Left != nil {
		root = root.Left
		return root
	}
	if root.Right != nil {
		root = root.Right
		return root
	}
	root = nil
	return root
}

func (t *Tree) min(root *Node) *Node {
	if root.Left == nil {
		return root
	}
	return t.min(root.Left)
}
