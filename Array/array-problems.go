package Array

//Two Sum Brute Force
func TwoSum(arr []int, target int) []int {
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i]+arr[j] == target {
				ans := make([]int, 2)
				ans[0] = i
				ans[1] = j
				return ans
			}
		}
	}

	return []int{}
}

//TwoSum Optimized
func TwoSumOptimized(nums []int, target int) []int {
	store := make(map[int]int) //map nums[i], index

	for i := 0; i < len(nums); i++ {
		val, ok := store[target-nums[i]]
		if ok {
			arr := make([]int, 2)
			arr[0] = val
			arr[1] = i
			return arr
		} else {
			store[nums[i]] = i
		}
	}

	return []int{}
}
