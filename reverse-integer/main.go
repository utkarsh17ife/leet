// https://leetcode.com/problems/reverse-integer/
package main

import "fmt"

func reverse(x int) int {
	var rev int
	keepLoop := true
	for keepLoop {
		le := x % 10
		x = x / 10
		if rev == 0 {
			rev = rev + le
		} else {
			rev = rev*10 + le
		}
		if x == 0 {
			keepLoop = false
		}
	}
	if -2147483648 > rev || rev > 2147483647 {
		return 0
	}
	return rev
}

func main() {

	fmt.Println(123, ": ", reverse(123))
	fmt.Println(-123, ": ", reverse(-123))
	fmt.Println(120, ": ", reverse(120))
	fmt.Println(0, ": ", reverse(0))
	fmt.Println(2147483147, ": ", reverse(2147483147))

}
