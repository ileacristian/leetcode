// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package add_two_numbers

func addNodeVals(l1 *ListNode, l2 *ListNode) (int, bool) {
	if l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		carry := false
		if sum >= 10 {
			sum = sum - 10
			carry = true
		}
		return sum, carry
	} else if l1 != nil && l2 == nil {
		return l1.Val, false
	} else if l1 == nil && l2 != nil {
		return l2.Val, false
	} else {
		return 0, false
	}
}

func getNextNode(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	val, carry := addNodeVals(l1, l2)

	result.Val = val
	currentNode := result
	if l1.Next != nil || l2.Next != nil || carry {
		result.Next = &ListNode{}
		currentNode = result.Next
	}

	for {
		l1 = getNextNode(l1)
		l2 = getNextNode(l2)

		if carry {
			currentNode.Val = 1
		}

		if l1 == nil && l2 == nil {
			break
		}

		val, carry = addNodeVals(l1, l2)
		currentNode.Val += val

		if currentNode.Val == 10 {
			currentNode.Val -= 10
			carry = true
		}

		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) || carry {
			nextNode := ListNode{}
			currentNode.Next = &nextNode
			currentNode = &nextNode
		}
	}

	return result
}
