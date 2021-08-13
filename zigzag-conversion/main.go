// https: //leetcode.com/problems/zigzag-conversion/
package main

import "fmt"

func convert(s string, numRows int) string {
	var rowStrMap = make(map[int]string)
	rowIndex := 1
	direction := ""

	if numRows == 1 {
		return s
	}

	for _, c := range s {
		val := rowStrMap[rowIndex]
		if val != "" {
			val = val + string(c)
			rowStrMap[rowIndex] = val
		} else {
			rowStrMap[rowIndex] = string(c)
		}
		if rowIndex == 1 || rowIndex == numRows {
			// change direction
			if direction == "down" {
				direction = "up"
			} else {
				direction = "down"
			}
		}
		if direction == "down" {
			rowIndex = rowIndex + 1
		} else {
			rowIndex = rowIndex - 1
		}
	}
	var finalString string
	// combine strings
	for i := 1; i <= numRows; i++ {
		finalString = finalString + rowStrMap[i]
	}
	return finalString
}

func main() {
	zigzagConverted := convert("ab", 1)
	fmt.Println(zigzagConverted)
}

// acegbdfh
