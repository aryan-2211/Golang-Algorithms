package Trees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (binaryTree *BinaryTree) PreOrderTraversal(root *TreeNode) {
	if root == nil {
		fmt.Println("the tree is empty")
		return
	}

	fmt.Printf("%d ", root.Val)
	binaryTree.PreOrderTraversal(root.Left)
	binaryTree.PreOrderTraversal(root.Right)
}

func (binaryTree *BinaryTree) InOrderTraversal(root *TreeNode) {
	if root == nil {
		fmt.Println("the tree is empty")
		return
	}

	binaryTree.InOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	binaryTree.InOrderTraversal(root.Right)
}

func (binaryTree *BinaryTree) PostOrderTraversal(root *TreeNode) {
	if root == nil {
		fmt.Println("the tree is empty")
		return
	}

	binaryTree.PostOrderTraversal(root.Left)
	binaryTree.PostOrderTraversal(root.Right)
	fmt.Printf("%d ", root.Val)
}
