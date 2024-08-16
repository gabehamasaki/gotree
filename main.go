package main

import "fmt"

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

func main() {
	array := [15]int64{
		34,
		84,
		15,
		0,
		2,
		99,
		79,
		9,
		88,
		89,
		18,
		31,
		39,
		100,
		101,
	}

	var root *Node

	for i := 0; i < 15; i++ {
		root = Insert(root, array[i])
		fmt.Printf("Inserted %d\n", array[i])
	}
	fmt.Println("InOrder traversal of the tree")
	InOrder(root)

	fmt.Println("\nSearching for 88")
	if Search(root, 88) != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}

func Insert(root *Node, data int64) *Node {
	temp := &Node{
		Value: data,
		Left:  nil,
		Right: nil,
	}

	if root == nil {
		root = temp
	} else {
		current := root
		var parent *Node

		for {
			parent = current
			if data < parent.Value {
				current = current.Right
				if current == nil {
					parent.Right = temp
					break
				}
			} else {
				current = current.Left
				if current == nil {
					parent.Left = temp
					break
				}
			}
		}
	}
	return root
}

func InOrder(root *Node) {
	if root != nil {
		InOrder(root.Right)
		fmt.Printf("%d ", root.Value)
		InOrder(root.Left)
	}
}

func Search(root *Node, data int64) *Node {
	if root == nil || root.Value == data {
		return root
	}

	if root.Value < data {
		return Search(root.Left, data)
	}
	return Search(root.Right, data)
}
