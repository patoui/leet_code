package main

import "fmt"

func strStr(haystack string, needle string) int {
    nlen := len(needle)
    hlen := len(haystack)

    if nlen > hlen {
        return -1
    }

    for i, v := range haystack {
        if v == rune(needle[0]) {
            j := 1
            for ; j < nlen; j++ {
                if i+j >= hlen || haystack[i+j] != needle[j] {
                    break
                }
            }

            if j == nlen {
                return i
            }
        }
    }

    return -1
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := strStr("sadbutsad", "sad")
    fmt.Printf("RESULT 1: %d\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := strStr("leetcode", "leeto")
    fmt.Printf("RESULT 2: %d\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := strStr("hello", "ll")
    fmt.Printf("RESULT 3: %d\n\n", r3)

    fmt.Println("EXAMPLE 4")
    r4 := strStr("mississippi", "issip")
    fmt.Printf("RESULT 4: %d\n\n", r4)

    fmt.Println("EXAMPLE 5")
    r5 := strStr("aaaa", "aaa")
    fmt.Printf("RESULT 5: %d\n\n", r5)
}
