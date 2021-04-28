package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

//tree
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

//node
func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		fmt.Printf("%d ", n.key)
		printPostOrder(n.right)
	}
}

func main() {
	var root Tree
	root.insert(50)
	root.insert(30)
	root.insert(20)
	root.insert(40)
	root.insert(70)
	root.insert(60)
	root.insert(80)
	printPostOrder(root.root)
}
