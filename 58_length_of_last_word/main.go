package main

import "fmt"

func lengthOfLastWord(s string) int {
    wl := 0

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != 32 {
            wl++
        } else if wl > 0 {
            return wl
        }
    }

    return wl
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := lengthOfLastWord("Hello World")
    fmt.Printf("RESULT 1: %d\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := lengthOfLastWord("   fly me   to   the moon  ")
    fmt.Printf("RESULT 2: %d\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := lengthOfLastWord("luffy is still joyboy")
    fmt.Printf("RESULT 3: %d\n\n", r3)

    fmt.Println("EXAMPLE 4")
    r4 := lengthOfLastWord("a")
    fmt.Printf("RESULT 4: %d\n\n", r4)
}
