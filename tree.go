// Package btree is simple binary tree implementation that can store
// any value type that implements ValueInterface.
// This implementation is a modified version of 'golang/tour/tree.go',
// (https://github.com/golang/tour/blob/master/tree/tree.go), adapted to
// store values of custom types.
package btree

import "fmt"

// ValueInterface should be implemented by the value type, which will be stored
// in the binary tree.
// See IntValue below for a sample implementation that stores a single integer
// value.
type ValueInterface interface {
	// Returns true if the interface has a smaller value than the argument passed
	// to the value parameter.
	Less(value interface{}) bool
}

// IntValue is sample implementaion of ValueInterface
type IntValue int

// Less returns true if the interface has a smaller value than the argument passed
// to the value parameter.
func (i IntValue) Less(value interface{}) bool {
	val := value.(IntValue)
	return int(i) < int(val)
}

// Node represent an element in the binary tree. Node could be the root node,
// an internal node or a leaf node.
// Left stores a pointer to a subtree with smaller values the current one.
// Right stores a pointer to a subtree with greater values the current one.
// Value should never be nil.
// Either or both of Left and Right could be nil.
type Node struct {
	Value ValueInterface
	Left  *Node
	Right *Node
}

// New creates a new node with the given value
func New(value ValueInterface) *Node {
	return &Node{Value: value}
}

// String method from https://github.com/golang/tour/blob/master/tree/tree.go
func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprint(n.Value)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}

// IsLeaf returns true if the TreeNode is a leaf in the tree (has no left or right children).
// A single root node is also considered a leaf.
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// Insert method from https://github.com/golang/tour/blob/master/tree/tree.go
func (n *Node) Insert(value ValueInterface) *Node {
	if n == nil {
		return &Node{value, nil, nil}
	}

	if value.Less(n.Value) {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}
	return n
}

// PrevLeaf returns the leaf node that is to the left of the specified node.
// If any of the children are internal nodes, traverses their children until
// a leaf is found.
func (n *Node) PrevLeaf() *Node {
	left := n.Left
	for !left.IsLeaf() {
		left = left.Right
	}
	return left
}

// NextLeaf returns the leaf node that is to the right of the specified node.
// If any of the children are internal nodes, traverses their children until
// a leaf is found.
func (n *Node) NextLeaf() *Node {
	right := n.Right
	for !right.IsLeaf() {
		right = right.Left
	}
	return right
}
