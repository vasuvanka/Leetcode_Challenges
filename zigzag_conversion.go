// The string "PAYPALISHIRING" is written in a zigzag pattern on a
// given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:

// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000

package main

import "fmt"

func convert(s string, numRows int) string {
    if numRows ==1 {
        return s
    }
    store := make(map[int]string)
    curRow := 0
    goingDown := false
    
    for start:=0; start < len(s); start++ {
        if val, ok := store[curRow]; ok {
            store[curRow]  = val + string(s[start])
        } else {
            store[curRow]  = string(s[start])
        }
        
        if curRow == 0 || curRow == numRows-1 {
            goingDown = !goingDown
        }
        if goingDown {
            curRow++
        } else {
            curRow--
        }
         
    }
    var result string
    for _,val := range store {
        result+=val
    }
    return result
    
}

func main()  {
	fmt.Println(convert("PAYPALISHIRING",3))
}
