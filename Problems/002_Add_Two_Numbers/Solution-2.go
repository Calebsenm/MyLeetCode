package main



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	A := l1
	B := l2
	var C *ListNode
	
	twoKeys := [2]bool{}
	sumTemp := 0
	tempA := 0
	tempB := 0
	numAdd := 0

	for {

		if A == nil {
			twoKeys[0] = true
			tempA = 0

		} else {
			tempA = A.Val
			A = A.Next
		}

		if B == nil {
			twoKeys[1] = true
			tempB = 0

		} else {
			tempB = B.Val
			B = B.Next
		}

        if twoKeys[0] == true && twoKeys[1] == true && numAdd == 0{
			break
		}

		sumTemp = tempA + tempB + numAdd

		if sumTemp >= 10 {
			sumTemp  = sumTemp - 10
			numAdd = 1
		} else {
			numAdd = 0
		}


		list := &ListNode{Val: sumTemp, Next: nil}

		if C == nil {
			C = list
		} else {

            current := C 

			for C.Next != nil {
				C = C.Next
			}

			C.Next = list
			C =  current	
		}
	}

	return C
}




