package main

import (
    "testing"
)

func BenchmarkStrStr(b *testing.B) {
    for i := 0; i < b.N; i++ {
        lengthOfLastWord("Hello World")
        lengthOfLastWord("   fly me   to   the moon  ")
        lengthOfLastWord("luffy is still joyboy")
    }
}
