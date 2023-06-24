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

    var ae string
    var be string
    ao := 1
    bo := 1
    c := ""
    o := ""

    for {
        av := ao > al
        bv := bo > bl

        if av && bv {
            break;
        }

        if av {
            ae = "0"
        } else {
            ae = string(a[al - ao])
            ao++
        }

        if bv {
            be = "0"
        } else {
            be = string(b[bl - bo])
            bo++
        }

        if ae == "1" && be == "1" {
            if c == "1" {
                o = "1" + o
            } else {
                o = "0" + o
                c = "1"
            }
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
