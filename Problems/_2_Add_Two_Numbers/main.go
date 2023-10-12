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

	theL1 := reverseLinkedList(l1)
	theL2 := reverseLinkedList(l2)


	for {
		newNode := &ListNode{}

		if theL1 == nil || theL2 == nil {
			break
		}

		if theL1.Val+theL2.Val+temp >= 10 {
			newNode = &ListNode{Val: theL1.Val+theL2.Val  -10 }
			temp = 1
		} else {

			newNode = &ListNode{Val:  theL1.Val + theL2.Val + temp}
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

		// 1 2 3 4
		theL1 = theL1.Next
		// 5 3 
		if theL2.Next == nil {
			theL2.Val = 0
		}	else {
			theL2 = theL2.Next
		}
		
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
