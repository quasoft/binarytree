package btree

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	want := IntValue(42)

	node := New(want)
	if node == nil {
		t.Fatalf("New() = nil, want a node struct")
	}

	got, ok := node.Value.(IntValue)
	if !ok {
		t.Fatalf("node.Value is not of type IntValue")
	}

	if got != want {
		t.Errorf("node.Value = %v, want %v", got, want)
	}
}

func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		t    *Node
		want string
	}{
		{"Empty", nil, "()"},
		{"Root only", &Node{IntValue(7), nil, nil}, "(7)"},
		{"Two levels", &Node{IntValue(7), &Node{IntValue(3), nil, nil}, &Node{IntValue(17), nil, nil}}, "((3) 7 (17))"},
		{"Three levels", &Node{IntValue(7), &Node{IntValue(3), &Node{IntValue(2), nil, nil}, nil}, &Node{IntValue(17), nil, nil}}, "(((2) 3) 7 (17))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_IsLeaf(t *testing.T) {
	//     7
	//   3    17
	// 2
	tree := &Node{IntValue(7), &Node{IntValue(3), &Node{IntValue(2), nil, nil}, nil}, &Node{IntValue(17), nil, nil}}

	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"Root (7)", tree, false},
		{"Internal node (3)", tree.Left, false},
		{"Left leaf (2)", tree.Left.Left, true},
		{"Right leaf (17)", tree.Right, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.IsLeaf(); got != tt.want {
				t.Errorf("With tree %v; %v.IsLeaf() = %v, want %v", tree, tt.node, got, tt.want)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	// Start with tree:
	//               7
	//       3               17
	//    2     6	    12         22
	//                           20  23
	//                         19
	tree := &Node{}
	tree.Value = IntValue(7)
	tree.Left = &Node{IntValue(3), &Node{IntValue(2), nil, nil}, &Node{IntValue(6), nil, nil}}
	tree.Right = &Node{IntValue(17), &Node{IntValue(12), nil, nil}, &Node{IntValue(22), nil, nil}}
	tree.Right.Right.Left = &Node{IntValue(20), &Node{IntValue(19), nil, nil}, nil}
	tree.Right.Right.Right = &Node{IntValue(23), nil, nil}

	value := 5
	tree.Insert(IntValue(value))
	want := "(((2) 3 ((5) 6)) 7 ((12) 17 (((19) 20) 22 (23))))"
	got := tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}

	value = 4
	tree.Insert(IntValue(value))
	want = "(((2) 3 (((4) 5) 6)) 7 ((12) 17 (((19) 20) 22 (23))))"
	got = tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}

	value = 18
	tree.Insert(IntValue(value))
	want = "(((2) 3 (((4) 5) 6)) 7 ((12) 17 ((((18) 19) 20) 22 (23))))"
	got = tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}
}

// Example of using the binary tree with the built-in IntValue type.
func Example() {
	tree := New(IntValue(42))
	tree.Insert(IntValue(24))
	tree.Insert(IntValue(3))
	tree.Insert(IntValue(43))

	fmt.Print(tree)

	// Output: (((3) 24) 42 (43))
}
