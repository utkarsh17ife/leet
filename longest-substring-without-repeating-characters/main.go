// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func lengthOfLongestSubstringApproach1(s string) int {

	isDistinct := func(s string) bool {
		visited := make(map[rune]bool)
		for _, c := range s {
			if ok := visited[c]; ok {
				return false
			}
			visited[c] = true
		}
		return true
	}

	maxLen := 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subStr := s[i : j+1]
			if isDistinct(subStr) {
				if len(subStr) > maxLen {
					maxLen = len(subStr)
				}
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstringApproach2(s string) int {

	// maintain a hash map for visited chars
	var i int
	visited := make(map[string]int)
	sos := 0
	maxSubStringLen := 0

	if len(s) == 0 {
		return 0
	}

	visited[string([]rune(s)[0])] = 0

	for i = 1; i < len(s); i++ {

		currChar := string([]rune(s)[i])
		var lastPos int

		if pos, ok := visited[currChar]; ok {
			lastPos = pos

			if lastPos < sos {

			} else {

				// calculate len of substring formed till this point
				subStrLen := i - sos
				if subStrLen > maxSubStringLen {
					maxSubStringLen = subStrLen
				}

				sos = lastPos + 1
			}

		}

		// update at last
		visited[currChar] = i
	}

	if maxSubStringLen < i-sos {
		maxSubStringLen = i - sos
	}

	return maxSubStringLen

}

func main() {
	fmt.Println("----Approach 1-----")
	fmt.Println("abcabcbb", lengthOfLongestSubstringApproach1("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstringApproach1("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstringApproach1("pwwkew"))
	fmt.Println(" ", lengthOfLongestSubstringApproach1(" "))
	fmt.Println("au", lengthOfLongestSubstringApproach1("au"))

	fmt.Println("----Approach 2-----")
	fmt.Println("abcabcbb", lengthOfLongestSubstringApproach2("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstringApproach2("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstringApproach2("pwwkew"))
	fmt.Println(" ", lengthOfLongestSubstringApproach2(" "))
	fmt.Println("au", lengthOfLongestSubstringApproach2("au"))
}

// explore
// https://www.geeksforgeeks.org/print-longest-substring-without-repeating-characters/

// loop over string

// set sos = 0
// max len = 0

// check the pos of the current char

// if last pos of current char < current sos
// update the pos

// if last pos of current char >= current sos

// len of sub string => curr - sos

// check and update max

// update the pos of chae

// reset sos to cur pos
