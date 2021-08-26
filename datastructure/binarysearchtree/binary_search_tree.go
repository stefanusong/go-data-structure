package binarysearchtree

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

// func createNode(value int) *Node {
// 	newNode := &Node{value, nil, nil}
// 	return newNode
// }

func (bst *BST) createNode(value int) {
	newNode := &Node{value, nil, nil}
	if bst.root == nil {
		bst.root = newNode
	} else {
		bst.root = insertNode(bst.root, newNode)
	}
}

func insertNode(root, newNode *Node) *Node {
	if newNode.value < root.value {
		root.left = insertNode(root.left, newNode)
	} else if newNode.value > root.value {
		root.right = insertNode(root.right, newNode)
	}
	return root
}

func (bst *BST) printAll() {
	
}

func Run() {

}
