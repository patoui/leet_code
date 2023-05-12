package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    lg, f, l := 0, 0, 0
    m := make(map[rune]int)
    for i, v := range(s) {
        pi, ok := m[v]
        if !ok || (i > pi && pi < f) {
            l = i
        } else {
            f, l = m[v]+1, m[v]+1
        }
        m[v] = i

        if l - f + 1 > lg {
            lg = l - f + 1
        }
    }

    return lg
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := lengthOfLongestSubstring("abcabcbb")
    fmt.Printf("RESULT 1: %d\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := lengthOfLongestSubstring("bbbbb")
    fmt.Printf("RESULT 2: %d\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := lengthOfLongestSubstring("pwwkew")
    fmt.Printf("RESULT 3: %d\n\n", r3)

    fmt.Println("EXAMPLE 4")
    r4 := lengthOfLongestSubstring("dvdf")
    fmt.Printf("RESULT 4: %d\n\n", r4)

    fmt.Println("EXAMPLE 5")
    r5 := lengthOfLongestSubstring("tmmzuxt")
    fmt.Printf("RESULT 5: %d\n\n", r5)
}
