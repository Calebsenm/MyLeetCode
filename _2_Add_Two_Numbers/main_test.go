package main

import (
	"testing"
	"fmt"
)
// This is a test for the function 

func TestTwoNumbers(t *testing.T){

	l1 := LinkedList{}
	l1.add(2)
	l1.add(4)
	l1.add(3)

	l2 := LinkedList{}
	l2.add(5)
	l2.add(6)
	l2.add(4)

	var data1 [] int = []int{ 2 , 4 , 3 }
	var data2 [] int = []int{ 5 , 6 , 4 }
	

	var data3 [] int = []int {7 , 0 , 8 } 
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

