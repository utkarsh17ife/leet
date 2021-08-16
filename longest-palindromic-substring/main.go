// https://leetcode.com/problems/longest-palindromic-substring/
package main

import "fmt"

func longestPalindrome(s string) string {

	start := 0
	end := 0

	expandFromMiddle := func(s string, l int, r int) int {
		if len(s) < 1 {
			return 0
		}
		for l > -1 && r < len(s) {
			if s[l] == s[r] {
				l--
				r++
			} else {
				break
			}
		}
		len := r - l - 1
		return len
	}

	for i, _ := range s {
		len1 := expandFromMiddle(s, i, i)
		len2 := expandFromMiddle(s, i, i+1)
		len := 0
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - ((len - 1) / 2)
			end = i + (len / 2)
		}
	}

	return s[start : end+1]

}

func main() {
	fmt.Println("werAAAAAAwjnnnwod : ", longestPalindrome("werAAAAAAwjnnnwod"))
	fmt.Println("cbbd : ", longestPalindrome("cbbd"))
	fmt.Println("bb : ", longestPalindrome("bb"))
	fmt.Println("a : ", longestPalindrome("a"))
	fmt.Println("aaa : ", longestPalindrome("aaa"))
	fmt.Println("aaaaaaaaba : ", longestPalindrome("aaaaaaaaba"))
	fmt.Println("tattarrattat : ", longestPalindrome("tattarrattat"))
}
