package hashfuncbench_test

import (
	"fmt"
	"strings"
	"testing"

	hashfuncbench "github.com/shunta0213/hash-func-bench"
)

var testSize = []int{10, 100, 1000, 10000}

func BenchmarkSHA256(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Sha256(key)
			}
		})
	}
}

func BenchmarkCRC32(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Crc32(key)
			}
		})
	}
}

func BenchmarkCRC64(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Crc64(key)
			}
		})
	}
}

func BenchmarkFNV(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Fnv(key)
			}
		})
	}
}

func BenchmarkMapHash(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Maphash(key)
			}
		})
	}
}
