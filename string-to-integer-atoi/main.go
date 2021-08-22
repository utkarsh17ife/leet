// https: //leetcode.com/problems/string-to-integer-atoi/
package main

//ASCII Values
// " " (white space): 32
// + (plus sign): 43
// - (minus sign): 45

import (
	"fmt"
	"strconv"
)

func myAtoi(s string) int {

	removeLeadingSpace := func(s string) string {
		spacesCount := 0
		for i, val := range s {
			if val == 32 {
				spacesCount = i + 1
				continue
			} else {
				break
			}
		}
		return s[spacesCount:]
	}

	readAndRemoveSign := func(s string) (string, string) {
		if s[0] == 45 {
			return "negetive", s[1:len(s)]
		}
		if s[0] == 43 {
			return "positive", s[1:len(s)]
		}
		return "positive", s
	}

	// read number till valid ASCII 48 - 57
	// remove any trailing extra characters
	readNumbers := func(s string) (bool, string) {
		zerosEndedYet := false
		numbersRead := 0
		startingZeros := 0
		for _, val := range s {
			if val == 48 && !zerosEndedYet {
				startingZeros++
				numbersRead++
			} else if val >= 48 && val <= 57 {
				zerosEndedYet = true
				numbersRead++
			} else {
				break
			}
		}
		if numbersRead > 0 {
			return true, s[startingZeros:numbersRead]
		} else {
			return false, ""
		}
	}

	// remove leading spaces
	s = removeLeadingSpace(s)

	if s == "" {
		return 0
	}

	sign, s := readAndRemoveSign(s)
	ok, s := readNumbers(s)
	if !ok {
		return 0
	}

	// convert number string to number
	num, _ := strconv.ParseInt(s, 0, 64)

	// add sign
	if sign == "negetive" {
		num = -1 * num
	}

	// handle number going out of the limits of int32
	if num > 2147483647 {
		return 2147483647
	} else if num < -2147483648 {
		return -2147483648
	}

	res := int(num)

	return res
}

func main() {
	fmt.Println("myAtoi(0):", myAtoi("0"))
	fmt.Println("myAtoi(    42):", myAtoi("    42"))
	fmt.Println("myAtoi(-000042):", myAtoi("-000042"))
	fmt.Println("myAtoi(    42):", myAtoi("   -42"))
	fmt.Println("myAtoi(4193 with words):", myAtoi("4193 with words"))
	fmt.Println("myAtoi(words and 987):", myAtoi("words and 987"))
	fmt.Println("myAtoi(-91283472332):", myAtoi("-91283472332"))
	fmt.Println("myAtoi(21474836460):", myAtoi("21474836460"))
	fmt.Println("myAtoi(  ):", myAtoi("  "))
}
