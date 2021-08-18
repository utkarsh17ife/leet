// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	nums3 := make([]int, 0)

	i := 0
	j := 0

	for i < lenNums1 || j < lenNums2 {

		if i >= lenNums1 && j >= lenNums2 {
			break
		}

		if i >= lenNums1 {
			nums3 = append(nums3, nums2[j])
			j++
			continue
		}

		if j >= lenNums2 {
			nums3 = append(nums3, nums1[i])
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			fmt.Println(3)
			nums3 = append(nums3, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			nums3 = append(nums3, nums2[j])
			j++
		} else if nums1[i] == nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		}

	}

	if len(nums3) == 0 {
		return float64(0)
	}
	if len(nums3) == 1 {
		return float64(nums3[0])
	}

	var median float64
	if len(nums3)%2 == 0 {
		median = float64(nums3[len(nums3)/2]+nums3[(len(nums3)/2)-1]) / 2
	} else {
		median = float64(nums3[len(nums3)/2])
	}

	return median

}

func main() {

	_ = findMedianSortedArrays([]int{2}, []int{})
}
