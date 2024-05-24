package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func(n *Node) Insert(k int) {

	// go to right
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}

	// go to left
	if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search
func(n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	}

	if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(40)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(53)
	tree.Insert(76)
	tree.Insert(26)
	tree.Insert(24)
	tree.Insert(123)
	
	fmt.Println(tree.Search(241))
}
