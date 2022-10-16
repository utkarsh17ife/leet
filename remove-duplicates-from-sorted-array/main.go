package main

import "fmt"

/*
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tempArr := []int{}
	for _, v := range nums {
		if len(tempArr) == 0 {
			tempArr = append(tempArr, v)
		} else if tempArr[len(tempArr)-1] != v {
			tempArr = append(tempArr, v)
		}
	}
	return copy(nums, tempArr)
}
*/

func removeDuplicates(nums []int) int {
	tmp := nums[0]
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > tmp {
			tmp = nums[i]
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 7, 8}
	k := removeDuplicates(arr)
	fmt.Println(k)
	fmt.Println(arr)
}
