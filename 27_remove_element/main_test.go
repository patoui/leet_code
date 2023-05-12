package main

import (
    "testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
    n := []int{0,1,2,2,3,0,4,2}

    for i := 0; i < b.N; i++ {
        removeElement(n, 2)
    }
}
