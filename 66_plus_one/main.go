package main

import "fmt"

func plusOne(digits []int) []int {
    r := true

    for i := len(digits)-1; i >= 0; i-- {
        if r {
            digits[i]++
            r = false
        }

        if digits[i] < 10 {
            break;
        }

        digits[i] -= 10
        r = true
    }

    if r {
        digits = append([]int{1}, digits...)
    }

    return digits
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := plusOne([]int{1,2,3})
    fmt.Printf("RESULT 1: %v\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := plusOne([]int{4,3,2,1})
    fmt.Printf("RESULT 2: %v\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := plusOne([]int{9})
    fmt.Printf("RESULT 3: %v\n\n", r3)

    fmt.Println("EXAMPLE 4")
    r4 := plusOne([]int{9,9})
    fmt.Printf("RESULT 4: %v\n\n", r4)

    fmt.Println("EXAMPLE 5")
    r5 := plusOne([]int{2,9,9})
    fmt.Printf("RESULT 5: %v\n\n", r5)
}
