// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order

// Example 1:

// Input: nums = [2,7,11,15], target = 9 [0:7,1:2]
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

package main

//TwoSum - will return indexs of elements sum equal to target
func TwoSum(nums []int,target int) []int  {
	store := make(map[int]int)
	for index, num := range nums {
		if  other,ok := store[target-num]; ok {
			return []int{other,index}
		}
		store[num] = index
	}
	return []int{}
}

// func main()  {
// 	fmt.Println(TwoSum([]int{2,7,11,15},13))
// }