package main

import (
	"fmt"
)
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	lenght int
}

func (l *LinkedList) add(value int) {

	newNode := new(ListNode)
	newNode.Val = value
	newNode.Next = nil

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.Next != nil; iterator = iterator.Next {
		}
	}

	l.lenght++
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	suma := new(ListNode)
	
	for {
		if l1 == nil || l2 == nil {
			break
		}


		suma.Val = l1.Val + l2.Val
		fmt.Println(suma , "XD")
		
		suma.Next = suma
		
		
		l1 = l1.Next
		l2 = l2.Next
	}
	return suma
}

func main() {
	
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

*/
