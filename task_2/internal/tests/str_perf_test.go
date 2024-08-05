package tests

import (
	"testing"
	"task/internal/stringers"
)

func BenchmarkReverseWordsFirstOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        stringers.ReverseWordsFirst("hello")
    }
}

func BenchmarkReverseWordsSecondOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        stringers.ReverseWordsSecond("hello")
    }
}