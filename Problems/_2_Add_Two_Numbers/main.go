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

func reverseLinkedList(list *ListNode) *ListNode {
	var reversed *ListNode
	toReverseNode := list
	newReversed := toReverseNode

	for {
	
		if toReverseNode.Next == nil {
			break
		}

		nexNode := toReverseNode.Next
		newNode := &ListNode{Val: toReverseNode.Next.Val}
		newReversed.Next = newNode
		newNode.Next = nexNode 

		if reversed == nil{
			reversed = newNode
		}else {
			reversed = newReversed
		}
	
		toReverseNode = toReverseNode.Next

	}

	return reversed
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	var current *ListNode

	for {
		newNode := &ListNode{}

		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val+l2.Val+temp >= 10 {
			temp = l1.Val + l2.Val + temp
			newNode = &ListNode{Val: 0}
		} else {

			newNode = &ListNode{Val: temp - l1.Val + l2.Val}
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

		l1 = l1.Next
		l2 = l2.Next

		//fmt.Println(current.Next)
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



for:

if a != null: 
	val = a + b
	if val >= 10:
    	rest = val -10
		val =  rest

if a == null: 
	val = a + b
	if val >= 10:
    	rest = val -10
		val =  rest

nodo -> = val 
*/
