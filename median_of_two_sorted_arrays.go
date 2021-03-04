// Median of Two Sorted Arrays

// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
// Example 3:

// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000
// Example 4:

// Input: nums1 = [], nums2 = [1]
// Output: 1.00000
// Example 5:

// Input: nums1 = [2], nums2 = []
// Output: 2.00000

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

package main

import "fmt"

func findMedian(list1,list2 []int) float32 {
	var combined []int
	combined = append(combined, list1...)
	combined = append(combined, list2...)
	fmt.Println(combined)
	var median float32
	if len(combined) < 1 {
		return median
	}
	result := MergeSort(combined)
	mid := len(result)/2
	if len(result) % 2 == 0 {
		median = float32(result[mid-1]+result[mid])/2
	} else {
		median = float32(result[mid])
	}
	return median
}
// MergeSort - merge sort
func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	// find mid
	mid := len(list)/2;
	return merge(MergeSort(list[0:mid]),MergeSort(list[mid:]))
}

func merge(left,right []int) []int {
	var result []int
	var leftIndex int
	var rightIndex int
	for  leftIndex < len(left) && rightIndex < len(right) {
		if(left[leftIndex]<right[rightIndex]){
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)
	return result
}

func main()  {
	fmt.Println(findMedian([]int{1,2},[]int{3,4,5}))
}