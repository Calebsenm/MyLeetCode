package main

import (
	"testing"
	"fmt"
)
// This is a test for the function 

func TestTwoNumbers(t *testing.T){

	l1 := LinkedList{}
	l1.add(9)
	l1.add(9)
	l1.add(4)

	l2 := LinkedList{}
	l2.add(1)
	l2.add(2)

	var data1 [] int = []int{ 9 , 9 , 4 }
	var data2 [] int = []int{     1 , 2 }
	

	var data3 [] int = []int {6 , 0 , 0 , 1 } 
	var data4 [4] int 

	nums := addTwoNumbers(l1.head , l2.head);

	var same bool  = true  
	
	i := 0
	for {
		//fmt.Println(nums)
		if nums == nil {
			break 
		}

		if data3[i] != nums.Val {
			same = false
		}

		data4[i] = nums.Val
		i++
		nums = nums.Next
	}
	
	if same == false {
		t.Errorf("Error La  Lista Esperada es: %d ", data3 )
	}	
	fmt.Println("Primera  " , data1 , "\nSegunda  ", data2 ,"\nEsperada ", data3, " \nResultado" , data4 )
		
}

func TestTwoNumbers2(t *testing.T){

	l1 := LinkedList{}
	l1.add(9)
	l1.add(9)
	l1.add(1)
	

	l2 := LinkedList{}
	l2.add(1)
	

	var data1 [] int = []int{ 9 , 9 , 1 }
	var data2 [] int = []int{ 1 }
	

	var data3 [] int = []int {0 , 0,  2 } 
	var data4 [3] int 

	nums := addTwoNumbers(l1.head , l2.head);

	var same bool  = true  
	
	i := 0
	for {
		//fmt.Println(nums)
		if nums == nil {
			break 
		}

		if data3[i] != nums.Val {
			same = false
		}

		data4[i] = nums.Val
		i++
		nums = nums.Next
	}
	
	if same == false {
		t.Errorf("Error La  Lista Esperada es: %d ", data3 )
	}else{
		fmt.Println("OKK ")
	}
	fmt.Println("Primera  " , data1 , "\nSegunda  ", data2 ,"\nEsperada ", data3, " \nResultado" , data4 )
		
}



func TestToReverse(t *testing.T){
	
	l1 := LinkedList{}
	l1.add(2)
	l1.add(9)
	l1.add(4)
	
	newNode := reverseLinkedList(l1.head)
	fmt.Println("")
	fmt.Println("-> ",newNode)
	fmt.Println("-> ",newNode.Next)
	fmt.Println("-> ",newNode.Next.Next)
	
}