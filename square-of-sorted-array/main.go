// https://leetcode.com/problems/squares-of-a-sorted-array/

package main

import "fmt"

func squareOfSortedArray(nums []int) []int {

	// init vales, set start and end pointers to 0 and len-1 resp.
	a := 0
	b := len(nums) - 1
	c := len(nums) - 1
	// result numsay
	resNums := make([]int, len(nums))

	// start loop from start and end
	for a < len(nums)-1 || b > -1 {

		// end loop
		if a > b {
			break
		}

		// compare the oppsite end values, based on bigger square value, push in result numsay from end
		if nums[a]*nums[a] > nums[b]*nums[b] {
			resNums[c] = nums[a] * nums[a]
			a++
		} else {
			resNums[c] = nums[b] * nums[b]
			b--
		}
		c--

	}

	return resNums
}

func main() {

	arr := []int{-7, -3, 2, 3, 11}
	arr = squareOfSortedArray(arr)
	fmt.Println(arr)

}
