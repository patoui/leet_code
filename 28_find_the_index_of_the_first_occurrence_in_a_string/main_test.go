package main

import (
    "testing"
)

func BenchmarkStrStr(b *testing.B) {
    for i := 0; i < b.N; i++ {
        strStr("sadbutsad", "sad")
        strStr("leetcode", "leeto")
        strStr("something", "me")
    }
}
