package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    o := head

    for head != nil {
        for head.Next != nil && head.Val == head.Next.Val {
            head.Next = head.Next.Next
        }

        head = head.Next
    }

    return o
}

func main() {
    one3 := ListNode{Val: 2}
    one2 := ListNode{Val: 1, Next: &one3}
    one1 := ListNode{Val: 1, Next: &one2}
    fmt.Println("EXAMPLE 1")
    r1 := deleteDuplicates(&one1)
    for r1 != nil {
        fmt.Printf("%d,", r1.Val)
        r1 = r1.Next
    }
    fmt.Printf("\n\n")

    two5 := ListNode{Val: 3}
    two4 := ListNode{Val: 3, Next: &two5}
    two3 := ListNode{Val: 2, Next: &two4}
    two2 := ListNode{Val: 1, Next: &two3}
    two1 := ListNode{Val: 1, Next: &two2}
    fmt.Println("EXAMPLE 1")
    r2 := deleteDuplicates(&two1)
    for r2 != nil {
        fmt.Printf("%d,", r2.Val)
        r2 = r2.Next
    }
    fmt.Printf("\n\n")

    three3 := ListNode{Val: 1}
    three2 := ListNode{Val: 1, Next: &three3}
    three1 := ListNode{Val: 1, Next: &three2}
    fmt.Println("EXAMPLE 1")
    r3 := deleteDuplicates(&three1)
    for r3 != nil {
        fmt.Printf("%d,", r3.Val)
        r3 = r3.Next
    }
    fmt.Printf("\n\n")
}
