package main

import (
	"fmt"

	"github.com/aryan/go-dsa/Array"
	"github.com/aryan/go-dsa/LinkedList"
	"github.com/aryan/go-dsa/Trees"
)

func main() {
	//for arrays
	arr := []int{1, 2, 3, 4, 5, 6}
	ans := make([]int, 2)
	ans = Array.TwoSumOptimized(arr, 9)
	fmt.Println(ans)

	//for linkedlist
	list := LinkedList.LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Display()

	//for Binary Trees
	node1 := &Trees.TreeNode{Val: 1}
	node2 := &Trees.TreeNode{Val: 2}
	node3 := &Trees.TreeNode{Val: 3}

	node1.Left = node2
	node1.Right = node3

	root := node1
	bt := Trees.BinaryTree{Root: node1}

	bt.PreOrderTraversal(root)
}
