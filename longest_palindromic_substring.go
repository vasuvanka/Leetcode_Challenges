// Longest Palindromic Substring

// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"
// Example 3:

// Input: s = "a"
// Output: "a"
// Example 4:

// Input: s = "ac"
// Output: "a"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters (lower-case and/or upper-case),

package main

import "fmt"


func longestPalindrome(s string) string {
    if len(s) < 2{
        return s
    }
    var start,end int
    for index:=0; index< len(s);index++{
        value1:= expandAround(s,index,index)
        value2:= expandAround(s,index,index+1)
        value:= value2
        if value1 > value2 {
            value = value1
        }
        if value > end-start {
            start = index-(value-1)/2
            end = index+value/2
        }
    }
    return string(s[start:end+1])
}

func expandAround(s string,left,right int) int {
    for left >= 0 && right< len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right-left-1
}

func main()  {
	fmt.Println(longestPalindrome("cbbd"))
}
