package main

import "fmt"

// see: https://medium.com/analytics-vidhya/leetcode-q70-climbing-stairs-easy-444a4aae54e8

func climbStairs(n int) int {
    m := make(map[int]int)
    return climb(n, m)
}

func climb(n int, m map[int]int) int {
    if v, ok := m[n]; ok {
        return v
    }

    if n == 1 || n == 2 {
        return n
    }

    r := climb(n - 1, m) + climb(n - 2, m)
    m[n] = r
    return r
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := climbStairs(2)
    fmt.Printf("RESULT 1: %v\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := climbStairs(3)
    fmt.Printf("RESULT 2: %v\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := climbStairs(4)
    fmt.Printf("RESULT 3: %v\n\n", r3)
}
