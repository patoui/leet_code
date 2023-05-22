package main

import "fmt"

func searchInsert(nums []int, target int) int {
    min := 0
    max := len(nums) - 1

    for min <= max {
        mid := (max + min) / 2

        if nums[mid] < target {
            min = mid + 1
        } else {
            max = mid - 1
        }
    }

    return min
}

func main() {
    fmt.Println("EXAMPLE 1")
    t1 := 5
    r1 := searchInsert([]int{1,3,5,6}, t1)
    fmt.Printf("TARGET 1: %d\n", t1)
    fmt.Printf("RESULT 1: %d\n\n", r1)

    fmt.Println("EXAMPLE 2")
    t2 := 2
    r2 := searchInsert([]int{1,3,5,6}, t2)
    fmt.Printf("TARGET 2: %d\n", t2)
    fmt.Printf("RESULT 2: %d\n\n", r2)

    fmt.Println("EXAMPLE 3")
    t3 := 7
    r3 := searchInsert([]int{1,3,5,6}, t3)
    fmt.Printf("TARGET 3: %d\n", t3)
    fmt.Printf("RESULT 3: %d\n\n", r3)

    fmt.Println("EXAMPLE 4")
    t4 := 7
    r4 := searchInsert([]int{1,3,4,5,6,7,9}, t4)
    fmt.Printf("TARGET 4: %d\n", t4)
    fmt.Printf("RESULT 4: %d\n\n", r4)

    fmt.Println("EXAMPLE 5")
    t5 := 0
    r5 := searchInsert([]int{1,3,5,6}, t5)
    fmt.Printf("TARGET 5: %d\n", t5)
    fmt.Printf("RESULT 5: %d\n\n", r5)

    fmt.Println("EXAMPLE 6")
    t6 := 2
    r6 := searchInsert([]int{1}, t6)
    fmt.Printf("TARGET 6: %d\n", t6)
    fmt.Printf("RESULT 6: %d\n\n", r6)

    fmt.Println("EXAMPLE 7")
    t7 := 1
    r7 := searchInsert([]int{1}, t7)
    fmt.Printf("TARGET 7: %d\n", t7)
    fmt.Printf("RESULT 7: %d\n\n", r7)

    fmt.Println("EXAMPLE 8")
    t8 := 2
    r8 := searchInsert([]int{1,3}, t8)
    fmt.Printf("TARGET 8: %d\n", t8)
    fmt.Printf("RESULT 8: %d\n\n", r8)
}
