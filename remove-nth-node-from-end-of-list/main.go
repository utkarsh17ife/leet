// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
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

// // slow and fast pointer approach
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	slow := head

	// move the fast pointer to the nth element
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// remove first element if fast is pointing to the last node
	if fast == nil {
		head = head.Next
		return head
	}

	// move both slow and fast pointer and when fast pointer reached the end, slow pointer will be at nth element from the last
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove the element to which slow is pointing
	slow.Next = slow.Next.Next

	return head
}

func main() {

	// make list 1
	head1 := &ListNode{}
	head1.Val = 1
	head1.append(2)
	head1.append(3)
	head1.append(4)
	head1.append(5)
	head1.append(6)
	head1.append(7)
	head1.print()
	newHead := removeNthFromEnd(head1, 2)
	newHead.print()
}
