package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// sum of node values
	sum := 0

	// multiplier, every time we traverse a node, increment by 10
	//	example: 
	//		iteration 1: 2 * 1 (multiplier) + 5 * 1 (multiplier)
	//		iteration 2: 4 * 10 (multiplier) + 6 * 10 (multiplier)
	//		iteration 3: 3 * 100 (multiplier) + 4 * 100 (multiplier)
	multiplier := 1
	
	for l1 != nil || l2 != nil {
		l1v := 0

		// if node exists, get its value
		if l1 != nil {
			l1v = l1.Val
		}

		l2v := 0

		// if node exists, get its value
		if l2 != nil {
			l2v = l2.Val
		}

		// add to sum
		sum += l1v * multiplier + l2v * multiplier

		// if node exist, get its next node
		if l1 != nil {
			l1 = l1.Next
		}
	
		// if node exist, get its next node
		if l2 != nil {
			l2 = l2.Next
		}

		multiplier *= 10
	}

	// modulus gives use the current node value as it's single digits (e.g. 0 to 9)
	n := ListNode{Val: sum % 10}
	// division by 10 and reassign to sum to move to the next value (e.g. 342 / 10 = 34)
	//	as golang removes the remainder from integer divisions
	sum /= 10

	// loop until the sum is larger than 10
	for sum > 0 {
		temp := n
		n = ListNode{Val: sum % 10, Next: &temp}
		sum /= 10
	}

	// vars to perform list reversal
	var curr *ListNode
	var prev *ListNode
	var next *ListNode
	curr = &n

	// reverse linked list
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr = prev

	return curr
}

func main() {
	a3 := ListNode{Val: 3}
	a2 := ListNode{Val: 4, Next: &a3}
	a1 := ListNode{Val: 2, Next: &a2}

	b3 := ListNode{Val: 4}
	b2 := ListNode{Val: 6, Next: &b3}
	b1 := ListNode{Val: 5, Next: &b2}

	n := addTwoNumbers(&a1, &b1)

	for n != nil {
		fmt.Printf("%+v\n", n)
		n = n.Next
	}

	c7 := ListNode{Val: 9}
	c6 := ListNode{Val: 9, Next: &c7}
	c5 := ListNode{Val: 9, Next: &c6}
	c4 := ListNode{Val: 9, Next: &c5}
	c3 := ListNode{Val: 9, Next: &c4}
	c2 := ListNode{Val: 9, Next: &c3}
	c1 := ListNode{Val: 9, Next: &c2}

	d4 := ListNode{Val: 9}
	d3 := ListNode{Val: 9, Next: &d4}
	d2 := ListNode{Val: 9, Next: &d3}
	d1 := ListNode{Val: 9, Next: &d2}

	n2 := addTwoNumbers(&c1, &d1)

	for n2 != nil {
		fmt.Printf("%+v\n", n2)
		n2 = n2.Next
	}
}

