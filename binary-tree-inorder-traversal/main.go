// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	arr := make([]int, 0)
	arr = append(arr, inorderTraversal(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)

	return arr
}

func main() {

	root := &TreeNode{40, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Left.Left = &TreeNode{10, nil, nil}
	root.Left.Left.Left = &TreeNode{5, nil, nil}
	root.Left.Right = &TreeNode{30, nil, nil}
	root.Right = &TreeNode{50, nil, nil}
	root.Right.Right = &TreeNode{60, nil, nil}
	root.Left.Right.Left = &TreeNode{67, nil, nil}
	root.Left.Right.Right = &TreeNode{78, nil, nil}

	fmt.Println(inorderTraversal(root))

}
