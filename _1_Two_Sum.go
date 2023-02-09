/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

*/

package main  

import "fmt"



func twoSum(numbers [] int , target int ) [] int{


  var newList [] int 

  for i:= 0; i < len( numbers ); i++{
    if numbers[i] < target{
      newList = append(numbers , numbers[i]);
    }
  }


  return newList
}

func main(){

  numbers := []int{2, 3, 5, 7, 11, 13}

  fmt.Println(twoSum(numbers , 9))
}










