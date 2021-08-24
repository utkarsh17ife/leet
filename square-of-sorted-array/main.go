package main

import "fmt"

func squareOfSortedArray(arr []int) []int {

	// init vales, set start and end pointers to 0 and len-1 resp.
	a := 0
	b := len(arr) - 1
	c := len(arr) - 1
	// result array
	resArr := make([]int, len(arr))

	// start loop from start and end
	for a < len(arr)-1 || b > -1 {

		// end loop
		if a > b {
			break
		}

		// compare the oppsite end values, based on bigger square value, push in result array from end
		if arr[a]*arr[a] > arr[b]*arr[b] {
			resArr[c] = arr[a] * arr[a]
			a++
		} else {
			resArr[c] = arr[b] * arr[b]
			b--
		}
		c--

	}

	return resArr
}

func main() {

	arr := []int{-7, -5, -3, -2, -1, 0, 1, 20}
	arr = squareOfSortedArray(arr)
	fmt.Println(arr)

}
