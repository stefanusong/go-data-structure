package binarysearchtree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func createNode(value int) *Node {
	newNode := &Node{value, nil, nil}
	// if bst.root == nil {
	// 	bst.root = newNode
	// } else {
	// 	insertNode(bst.root, newNode)
	// }
	return newNode
}

func insertNode(root *Node, value int) *Node {
	// if newNode.value < root.value {
	// 	// Insert to left child if not exists
	// 	if root.left == nil {
	// 		root.left = newNode
	// 	} else { // else recursively check the left subtree
	// 		insertNode(root.left, newNode)
	// 	}
	// } else {
	// 	// Insert to right child if not exists
	// 	if root.left == nil {
	// 		root.left = newNode
	// 	} else { // else recursively check the right subtree
	// 		insertNode(root.left, newNode)
	// 	}
	// }
	if root == nil {
		return createNode(value)
	} else if value < root.value {
		root.left = insertNode(root.left, value)
	} else if value > root.value {
		root.right = insertNode(root.right, value)
	}
	return root
}

func removeNode(root *Node, value int) *Node {
	if root == nil {
		return root
	} else if value < root.value {
		root.left = removeNode(root.left, value)
	} else if value > root.value {
		root.right = removeNode(root.right, value)
	} else {
		// The to be deleted node found !

		// Check if this node has only one child that will replace its position
		if root.left == nil || root.right == nil {
			var replacingNode *Node
			// Set replacing node to the node that exists
			if root.left != nil {
				replacingNode = root.left
			} else {
				replacingNode = root.right
			}
			root.left = nil
			root.right = nil
			root = nil
			return replacingNode
		} else { // If this node has 2 child
			newRoot := Predecessor(root)
			root.value = newRoot.value
			root.left = removeNode(root.left, newRoot.value) //fix the left subtree
		}
	}
	return root
}

func Predecessor(root *Node) *Node {
	currentNode := root.left

	// Find highest value in the left subtree by traversing to the right of the tree
	for currentNode != nil && currentNode.right != nil {
		currentNode = currentNode.right
	}

	return currentNode
}

func searchNode(root *Node, value int) bool {

	if root != nil {
		if root.value == value {
			return true
		} else if value < root.value {
			return searchNode(root.left, value)
		} else if value > root.value {
			return searchNode(root.right, value)
		}
	}

	return false
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Print(root.value, "-> ")
		preOrder(root.left)
		preOrder(root.right)
	}
}

func inOrder(root *Node) {
	if root != nil {
		inOrder(root.left)
		fmt.Print(root.value, "-> ")
		inOrder(root.right)
	}
}

func postOrder(root *Node) {
	if root != nil {
		postOrder(root.left)
		postOrder(root.right)
		fmt.Print(root.value, "-> ")
	}
}

func Run() {
	bst := &BST{}
	items := []int{13, 12, 16, 15, 8, 14, 11, 18}
	totalItem := len(items)

	for i := 0; i < totalItem; i++ {
		bst.root = insertNode(bst.root, items[i])
		fmt.Println("Inserted", items[i])
		inOrder(bst.root)
		fmt.Println("")
	}

	for i := 0; i < totalItem; i++ {
		bst.root = removeNode(bst.root, items[i])
		fmt.Println("Removed", items[i])
		inOrder(bst.root)
		fmt.Println("")
	}

}
