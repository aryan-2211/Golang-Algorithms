package main

import (
	"fmt"

	"github.com/aryan/go-dsa/Array"
	"github.com/aryan/go-dsa/LinkedList"
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
}
