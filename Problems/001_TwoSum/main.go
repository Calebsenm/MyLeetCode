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

//funcion  con unbucle  que  recorre 2  veces  la lista 
//el  segundo empieza en la  iteracion delsegundo
// 1 2 3  primera 
// 2 3    segunda
// 3      tercera

package main

import "fmt"

func twoSum(nums []int, target int) [][]int {
	
	var newList [][]int
  
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {

			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				newList = append(newList, []int{i, j})
			}
		}

	}
	return newList
}

func main() {
	numbers := []int{-2, -2}
	fmt.Println(twoSum(numbers, -4))
}

