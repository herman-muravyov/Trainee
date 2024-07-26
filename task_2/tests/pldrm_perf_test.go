package tests

import (
    "testing"
    "task/palindrome"
)

func BenchmarkPalindromeFirstOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        palindrome.IsPalindromeFirst("anna")
    }
}

func BenchmarkPalindromeSecondOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        palindrome.IsPalindromeSecond("anna")
    }
}

func BenchmarkPalindromeThirdOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        palindrome.IsPalindromeThird("anna")
    }
}