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
    }

    var ae string
    var be string
    carry := ""
    o := ""

    for i := al - 1; i >= 0; i-- {
        ae = string(a[i])
        be = string(b[i])

        if carry == "1" && ae == "1" && be == "1" {
            o = "1" + o
        } else if ae == "1" && be == "1" {
            o = "0" + o
            carry = "1"
        } else if ae == "1" || be == "1" {
            if carry == "1" {
                o = "0" + o
            } else {
                o = "1" + o
            }
        } else if carry == "1" {
            o = "1" + o
            carry = "0"
        } else {
            o = "0" + o
        }
    }

    if carry == "1" {
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
}
