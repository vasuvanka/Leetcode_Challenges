package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	left := 0;
	right:= len(s)-1;

	for s[left] == s[right] {
		if left >= right {
			return true
		}
		left++
		right--
	}
	return false
}

func main(){
	fmt.Println(isPalindrome("racecar"))
}