# binarytree
Package binarytree is simple binary tree implementation that can store
any value type that implements ValueInterface.

This implementation is a modified version of [golang/tour/tree.go](https://github.com/golang/tour/blob/master/tree/tree.go), adapted to store values of custom types.

Currently does *NOT* support balancing.

Initially created for use in a voronoi diagram generator.

# How to use:

```go
import "github.com/quasoft/binarytree"

tree := binarytree.New(IntValue(42))
tree.Insert(IntValue(24))
tree.Insert(IntValue(3))
tree.Insert(IntValue(43))

fmt.Print(tree)

// Output: (((3) 24) 42 (43))
```