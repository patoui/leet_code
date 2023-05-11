package main

import (
    "testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
    n2 := []int{0,0,1,1,1,2,2,3,3,4}

    for i := 0; i < b.N; i++ {
        removeDuplicates(n2)
    }
}
