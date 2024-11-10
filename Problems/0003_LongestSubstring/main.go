package main 

import (
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
	var sum int
	var words string = s 
    var lenWords int = len(words)

	if lenWords == 0 {
		sum = 0
	} else if lenWords == 1 {
		sum = 1

	} else {

		punt := 0
		isDiferent := false
		sumT := 0

		for i := 1; i < lenWords ; i++ {

			if words[punt] == words[i] {
				punt += 1

			} else {
				for j := i; j > punt; j-- {
					if words[i] != words[j-1] {
						isDiferent = true

					} else {
						isDiferent = false
						sumT = 0
						i = punt +1
                        punt += 1
						break
					}
				}

				if isDiferent {
					sumT++
					if sumT >= sum {
						sum = sumT + 1
					}
				}

			}

		}

	}

	if lenWords != 0 && sum == 0 {
		sum += 1
	}

	return sum
}


func main(){
    var substring string = "abcabcbb";
    var num int = lengthOfLongestSubstring(substring);
    
    fmt.Println(num);
}
