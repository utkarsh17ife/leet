// https: //leetcode.com/problems/swapping-nodes-in-a-linked-list/

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

func swapNodes(head *ListNode, k int) *ListNode {

	findKthNodeFromLast := func(head *ListNode, k int) *ListNode {

		fast := head
		slow := head

		kthFrmLast := head

		for i := 0; i < k; i++ {
			fast = fast.Next
		}

		if fast == nil {
			return kthFrmLast
		}

		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}

		kthFrmLast = slow.Next

		return kthFrmLast
	}

	findKthNodeFromStart := func(head *ListNode, k int) *ListNode {

		currNode := head

		for i := 1; i < k; i++ {
			currNode = currNode.Next
		}

		return currNode
	}

	// find nodes
	kthNodeFromLast := findKthNodeFromLast(head, k)
	kthNodeFromStart := findKthNodeFromStart(head, k)

	// swap values
	temp := kthNodeFromLast.Val
	kthNodeFromLast.Val = kthNodeFromStart.Val
	kthNodeFromStart.Val = temp

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
	head1.append(8)
	head1.append(9)
	head1.print()
	swapNodes(head1, 1)
	head1.print()
	swapNodes(head1, 3)
	head1.print()
	swapNodes(head1, 1)
	head1.print()
	swapNodes(head1, 8)
	head1.print()

}
