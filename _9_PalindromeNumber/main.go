package main

import (
    "fmt"
    "strconv"
)

func isPalindrome(x int ) bool{

    num  := strconv.Itoa(x);

	i := 0
	j :=  len(num) - 1

	for i < len(num){

		if num[i] != num[j] {
			return false
		}
		i++
		j--
	}
    return true 
}


func main(){
    fmt.Println( isPalindrome(12321) );
}

