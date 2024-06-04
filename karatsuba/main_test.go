package main

import (
	"testing"
)

func BenchmarkKaratsuba_1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        karatsuba(1222, 2022)        
    }
}

func BenchmarkBruceforce_1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        multiply(1222, 2022)        
    }
}

func BenchmarkKaratsuba_2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        karatsuba(123456789, 987654321)        
    }
}

func BenchmarkBruceforce_2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        multiply(123456789, 987654321)        
    }
}
