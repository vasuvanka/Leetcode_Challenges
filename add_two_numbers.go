// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
//  Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]
// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// Constraints:

// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode - list node
type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	carry := 0
	var head ListNode
	current := &head
	for a != nil || b != nil || carry > 0{
		sum := carry
		if a != nil {
			sum+=a.Val
		} 
		if b != nil {
			sum+=b.Val
		}
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		current.Next = node
		current = current.Next
		if a.Next != nil {
			a = a.Next
		}
		if b.Next != nil {
			b = b.Next
		}
	}
	return head.Next
}