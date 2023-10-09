package main


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
		temp := l.head
		for {
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		temp.Next = newNode
	}

	l.lenght++
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	var sumList *ListNode

	for {
		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val+l2.Val+temp == 10 {
			temp = 1
			sumList = &ListNode{0, sumList}
		} else {
			sumList = &ListNode{l1.Val + l2.Val + temp,sumList}
			temp = 0
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	
	return sumList
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
