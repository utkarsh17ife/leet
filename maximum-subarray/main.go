// https://leetcode.com/problems/maximum-subarray/

package main

import "fmt"

func maxSubArray(nums []int) int {

	var max, currMax int
	max = nums[0]
	currMax = nums[0]

	for i := 1; i < len(nums); i++ {

		currMax = currMax + nums[i]
		if currMax < nums[i] {
			currMax = nums[i]
		}

		if currMax > max {
			max = currMax
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
