package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	multiplier := 1
	
	for {
		l1v := 0
		if l1 != nil {
			l1v = l1.Val
		}

		l2v := 0
		if l2 != nil {
			l2v = l2.Val
		}

		sum += l1v * multiplier + l2v * multiplier

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break;
		}

		multiplier *= 10
	}

	n := ListNode{Val: sum % 10}
	sum /= 10

	for sum > 0 {
		temp := n
		n = ListNode{Val: sum % 10, Next: &temp}
		sum /= 10
	}

	var curr *ListNode
	var prev *ListNode
	var next *ListNode
	curr = &n

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
	// a3 := ListNode{Val: 3}
	// a2 := ListNode{Val: 4, Next: &a3}
	// a1 := ListNode{Val: 2, Next: &a2}
	//
	// b3 := ListNode{Val: 4}
	// b2 := ListNode{Val: 6, Next: &b3}
	// b1 := ListNode{Val: 5, Next: &b2}
	//
	// n := addTwoNumbers(&a1, &b1)
	//
	// for n != nil {
	// 	fmt.Printf("%+v\n", n)
	// 	n = n.Next
	// }

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

