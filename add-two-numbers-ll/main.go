// https://leetcode.com/problems/add-two-numbers/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) append(val int) {
	// node last node
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	curHead := ln
	for curHead.Next != nil {
		curHead = curHead.Next
	}
	curHead.Next = newNode
	return
}

// reverse linked list
func (ln *ListNode) reverse() *ListNode {
	curr := ln
	var prev *ListNode
	var next *ListNode
	for curr != nil {
		// reverse
		next = curr.Next
		curr.Next = prev

		// jump ahead, prepare for next reverse
		prev = curr
		curr = next
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result, head *ListNode

	carry := 0

	for l1 != nil || l2 != nil {

		sum := 0

		if l1 != nil {
			sum = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		sum = sum + carry

		value := sum % 10
		carry = sum / 10

		// create node
		node := &ListNode{
			Val: value,
		}

		if result != nil {
			result.Next = node
			result = result.Next
		} else {
			result = node
			head = node
		}

	}

	if carry > 0 {
		// create node
		node := &ListNode{
			Val: carry,
		}
		result.Next = node
	}

	return head
}

func addTwoNumbersFaster(l1 *ListNode, l2 *ListNode) *ListNode {

	var result, head *ListNode

	carry := 0

	for l1 != nil || l2 != nil {

		sum := 0

		if l1 != nil {
			sum = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		sum = sum + carry

		value := sum % 10
		carry = sum / 10

		// create node
		node := &ListNode{
			Val: value,
		}

		if result != nil {
			result.Next = node
			result = result.Next
		} else {
			result = node
			head = node
		}

	}

	if carry > 0 {
		// create node
		node := &ListNode{
			Val: carry,
		}
		result.Next = node
	}

	return head
}

func (ln ListNode) print() {
	curHead := &ln
	for curHead.Next != nil {
		fmt.Print(curHead.Val)
		curHead = curHead.Next
	}
	fmt.Println(curHead.Val)
	return
}

func main() {
	// make list 1
	head1 := &ListNode{}
	head1.Val = 2
	head1.append(4)
	head1.append(9)

	// make list 2
	head2 := &ListNode{}
	head2.Val = 5
	head2.append(6)
	head2.append(4)
	head2.append(9)

	fmt.Println("first linked list")
	head1.print()
	fmt.Println("second linked list")
	head2.print()

	sumHead := addTwoNumbers(head1, head2)

	sumHead.print()

}
