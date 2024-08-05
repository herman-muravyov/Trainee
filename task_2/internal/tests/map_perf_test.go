package tests

import (
    "testing"
    "task/internal/mapping"
    "task/pkg/variables"
)

func BenchmarkMapFirstOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mapping.Mapping(variables.Slice)
    }
}

func BenchmarkSolutionSecondOptimization(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mapping.RemoveDuplicates(variables.Slice)
    }
}