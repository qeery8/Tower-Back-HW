package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func newNode(v int) *TreeNode {
	return &TreeNode{data: v, left: nil, right: nil}
}

type BST struct {
	root *TreeNode
}

func (bst *BST) Add(v int) {
	bst.root = add(bst.root, v)
}

func add(node *TreeNode, v int) *TreeNode {
	if node == nil {
		return newNode(v)
	}

	if v < node.data {
		node.left = add(node.left, v)
	} else {
		node.right = add(node.right, v)
	}
	return node
}

func (bst *BST) Delete(v int) {
	bst.root = delete(bst.root, v)
}

func delete(node *TreeNode, v int) *TreeNode {
	if node == nil {
		return nil
	}

	if v < node.data {
		node.left = delete(node.left, v)
	} else if v > node.data {
		node.right = delete(node.right, v)
	} else {
		if node.left == nil {
			temp := node.right
			node = nil
			return temp
		} else if node.right == nil {
			temp := node.left
			node = nil
			return temp
		}

		temp := minNode(node.right)
		node.data = temp.data
		node.right = delete(node.right, temp.data)
	}
	return node
}

func minNode(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BST) IsExist(v int) bool {
	return isExist(bst.root, v)
}

func isExist(node *TreeNode, v int) bool {
	if node == nil {
		return false
	}

	if v == node.data {
		return true
	} else if v < node.data {
		return isExist(node.left, v)
	} else {
		return isExist(node.right, v)
	}
}

func main() {
	bst := BST{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(1)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)
	bst.Add(9)

	if bst.IsExist(9) {
		fmt.Println("Найдено 9")
	} else {
		fmt.Println("9 - не найдено")
	}

	if bst.IsExist(4) {
		fmt.Println("Найдено 4")
	} else {
		fmt.Println("4 - не найдено")
	}

	bst.Delete(4)

	if bst.IsExist(4) {
		fmt.Println("Найдено 4")
	} else {
		fmt.Println("4 - не найдено")
	}

}
