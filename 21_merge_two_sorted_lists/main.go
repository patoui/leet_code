package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cn := &ListNode{Val: -100}
	var head *ListNode

	for list1 != nil || list2 != nil {
		if (list1 != nil && list2 != nil) {
			if (list1.Val < list2.Val) {
				cn.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				cn.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
		} else if (list1 != nil) {
			cn.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cn.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		cn = cn.Next

		if head == nil {
			head = cn
		}
	}

	return head
}

func main() {
	fmt.Println("EXAMPLE 1")

	a3 := ListNode{Val: 4}
	a2 := ListNode{Val: 2, Next: &a3}
	a1 := ListNode{Val: 1, Next: &a2}

	b3 := ListNode{Val: 4}
	b2 := ListNode{Val: 3, Next: &b3}
	b1 := ListNode{Val: 1, Next: &b2}

	result1 := mergeTwoLists(&a1,&b1)
	for {
		fmt.Printf("NODE VAL: %d\n", result1.Val)
		result1 = result1.Next
		if (result1 == nil) {
			break;
		}
	}

	fmt.Println("EXAMPLE 2")

	c1 := ListNode{Val: 0}

	result2 := mergeTwoLists(&c1,nil)
	for {
		fmt.Printf("NODE VAL: %d\n", result2.Val)
		result2 = result2.Next
		if (result2 == nil) {
			break;
		}
	}
}
