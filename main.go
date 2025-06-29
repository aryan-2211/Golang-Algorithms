package main

import (
	"fmt"

	"github.com/aryan/go-dsa/Array"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ans := make([]int, 2)

	ans = Array.TwoSumOptimized(arr, 9)
	fmt.Println(ans)
}
