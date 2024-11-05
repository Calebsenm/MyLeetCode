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

// This function is for reverse the linkedList
func reverseLinkedList(lst *ListNode) *ListNode {
	var reversed *ListNode
	theList := lst

	for {
		if theList == nil {
			break
		}

		newNode := &ListNode{Val: theList.Val}

		if reversed == nil {
			reversed = newNode
		} else {
			current := reversed
			newN := newNode
			newN.Next = current
			reversed = newN
		}

		theList = theList.Next
	}

	return reversed
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	var current *ListNode

	if l1.Next == nil && l1.Val == 0{
		current = l2
	}else{



	newl1 := l1
	newl2 := l2

	for {

		if newl1.Next == nil {
			newl1 = l2
			newl2 = l1
			break
		} else if newl2.Next == nil {
			newl1 = reverseLinkedList(l1)
			newl2 = reverseLinkedList(l2)
			break 

		}else{
			break
		}

		//newl1 = newl1.Next
		//newl2 = newl2.Next
	}

	theL1 := newl1
	theL2 := newl2

	if newl2.Next == nil {
		theL1 = reverseLinkedList(newl1)
		theL2 = reverseLinkedList(newl2)
	}else {
		theL2 = newl2
		theL1 = newl1
	}

	if l1.Next == nil{
		theL2 = l2 
		theL1 = l1
	}	

	for {
		newNode := &ListNode{}

		if theL1 == nil || theL2 == nil {
			break
		}

		sum := theL1.Val + theL2.Val
		if sum+temp >= 10 {
			newNode = &ListNode{Val: sum + temp - 10}
			temp = 1
		} else {
			newNode = &ListNode{Val: sum + temp}
			temp = 0
		}

		if current == nil {
			current = newNode

		} else {

			current1 := current

			for current.Next != nil {
				current = current.Next
			}
			current.Next = newNode
			current = current1
		}

		if theL1.Next == nil {
			theL1.Val = 0
		} else {
			theL1 = theL1.Next
		}

		if theL2.Next == nil {
			theL2.Val = 0
		} else {
			theL2 = theL2.Next
		}

		if theL1.Next == nil && theL1.Val == 0 && theL2.Next == nil && theL2.Val == 0 {
			break
		}
		//fmt.Println(current.Next)
	}

	if temp == 1 {
		theLastest := &ListNode{Val: temp}
		current1 := current

		for current.Next != nil {
			current = current.Next
		}
		current.Next = theLastest
		current = current1

	}
	}
	return current
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

/*
Solucion

a = 2 -> 9 -> 4
b = 1 -> 2

2 -> 9 -> 4
4 -> 9 -> 2




*/
