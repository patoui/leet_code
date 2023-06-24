package main

import "fmt"

func addBinary(a string, b string) string {
    al := len(a)
    bl := len(b)

    if al == 0 {
        return b
    }

    if bl == 0 {
        return a
    }

    if al > bl {
        for i := 0; i < al - bl; i++ {
            b = "0" + b
        }
        bl = al
    } else if bl > al {
        for i := 0; i < bl - al; i++ {
            a = "0" + a
        }
        al = bl
    }

    var ae string
    var be string
    c := ""
    o := ""

    for i := al - 1; i >= 0; i-- {
        ae = string(a[i])
        be = string(b[i])

        if c == "1" && ae == "1" && be == "1" {
            o = "1" + o
        } else if ae == "1" && be == "1" {
            o = "0" + o
            c = "1"
        } else if ae == "1" || be == "1" {
            if c == "1" {
                o = "0" + o
            } else {
                o = "1" + o
            }
        } else if c == "1" {
            o = "1" + o
            c = "0"
        } else {
            o = "0" + o
        }
    }

    if c == "1" {
        o = "1" + o
    }

    return o
}

func main() {
    fmt.Println("EXAMPLE 1")
    r1 := addBinary("11", "1")
    fmt.Printf("RESULT 1: %v\n\n", r1)

    fmt.Println("EXAMPLE 2")
    r2 := addBinary("1010", "1011")
    fmt.Printf("RESULT 2: %v\n\n", r2)

    fmt.Println("EXAMPLE 3")
    r3 := addBinary("1", "111")
    fmt.Printf("RESULT 3: %v\n\n", r3)
}
