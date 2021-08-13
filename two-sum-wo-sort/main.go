// https://leetcode.com/problems/two-sum/ (for non sorted array)
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	// make hash table
	ht := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ht[nums[i]] = i
	}

	a := -1
	b := -1

	// start from i and look for other number
	for i := 0; i < len(nums)-1; i++ {

		otherNum := target - nums[i]

		if j, ok := ht[otherNum]; ok {
			if i != j {
				a = i
				b = j
				break
			}
		}

	}

	return []int{a, b}

}

func main() {

	nums := []int{3, 2, 4}

	res := twoSum(nums, 6)

	fmt.Println(res[0], res[1])
}
