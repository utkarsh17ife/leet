// https://leetcode.com/problems/two-sum/ (for sorted array)
package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {

	i := 0
	j := len(nums) - 1

	for i < j {

		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] == target {
			break
		}

	}

	return []int{i + 1, j + 1}
}

func main() {

	nums := []int{1, 2, 4, 8, 5, 6, 86, 13, 71, 111, 51}
	sort.Ints(nums)
	res := twoSum(nums, 59)

	fmt.Println(res[0], res[1])
}
