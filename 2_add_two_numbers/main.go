package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// track head node
	head := &ListNode{}
	c := head

	// variables to track across loops
	// track previous node
	var p *ListNode
	// track carry over (e.g. 9+9 = 18, current node value = 8, carry over = 1)
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		l1v := 0
		// if l1 has a node available, retrieve it's value
		if l1 != nil {
			l1v = l1.Val
		}

		l2v := 0
		// if l2 has a node available, retrieve it's value
		if l2 != nil {
			l2v = l2.Val
		}

		cVal := l1v + l2v + carry

		// if the current value is above 9, get it's remainder and set the carry over
		if cVal > 9 {
			carry = 1
			cVal = cVal % 10
		} else {
			carry = 0
		}

		c.Val = cVal

		// if there is a previous node, set it's next node to the current node (list reversal)
		if p != nil {
			p.Next = c
		}

		// move to the next node if l1 has a current node
		if l1 != nil {
			l1 = l1.Next
		}

		// move to the next node if l2 has a current node
		if l2 != nil {
			l2 = l2.Next
		}

		// set the previous node to the current node
		p = c
		// initialize a new current node
		c = &ListNode{}
	}

	return head
}

func main() {
	fmt.Println("EXAMPLE 1")
	a3 := ListNode{Val: 3}
	a2 := ListNode{Val: 4, Next: &a3}
	a1 := ListNode{Val: 2, Next: &a2}

	b3 := ListNode{Val: 4}
	b2 := ListNode{Val: 6, Next: &b3}
	b1 := ListNode{Val: 5, Next: &b2}

	n := addTwoNumbers(&a1, &b1)

	for n != nil {
		fmt.Printf("%d,", n.Val)
		n = n.Next
	}

	fmt.Println()

	fmt.Println("EXAMPLE 2")
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
		fmt.Printf("%d,", n2.Val)
		n2 = n2.Next
	}
}

