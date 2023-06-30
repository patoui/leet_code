package main

import "fmt"

func mySqrt(x int) int {
    i := 1
    for {
        if i * i > x {
            return i-1
        }

        i++
    }
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := mySqrt(4)
    fmt.Printf("RESULT 1: %v\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := mySqrt(8)
    fmt.Printf("RESULT 2: %v\n\n", r2)
}
