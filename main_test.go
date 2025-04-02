package main

import "testing"

func BenchmarkSilly(b *testing.B) {
	var sillyResult float64
	for i := 0; i < b.N; i++ {
		sillyResult += float64(i)
		sillyResult *= float64(i)
		divisor := float64(i) + 1
		if divisor > 0 {
			sillyResult /= divisor
		}
	}
	_ = sillyResult // to avoid compiler optimization
}
