package main

import (
	"fmt"
)

func parnToNum(paran string) int32 {
	switch paran {
	case "[":
		return 1
	case "(":
		return 2
	case "{":
		return 3
	case "]":
		return -1
	case ")":
		return -2
	case "}":
		return -3
	default:
		return 0
	}
}

func isValid(s string) bool {
	arr := []int32{}
	isValid := true
	for i := 0; i < len(s); i++ {
		c := parnToNum(string(s[i]))
		if c < 0 {
			if len(arr) == 0 {
				isValid = false
				break
			}
			lastElem := arr[len(arr)-1]
			if c+lastElem == 0 {
				arr = arr[:len(arr)-1]
			} else {
				isValid = false
				break
			}
		} else {
			arr = append(arr, c)
		}
	}
	if len(arr) != 0 {
		isValid = false
	}
	return isValid
}

func main() {
	str := "{[{()}]}"
	resp := isValid(str)
	fmt.Println(resp)
}
