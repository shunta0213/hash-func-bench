package hashfuncbench_test

import (
	"fmt"
	"strings"
	"testing"

	hashfuncbench "github.com/shunta0213/hash-func-bench"
)

var testSize = make([]int, 200)

func init() {
	for i := 0; i < 200; i++ {
		testSize[i] = i + 1
	}
}

func BenchmarkSHA256(b *testing.B) {
	for _, size := range testSize {
		b.Run(fmt.Sprintf("SHA256_%d", size), func(b *testing.B) {
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
		b.Run(fmt.Sprintf("CRC32_%d", size), func(b *testing.B) {
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
		b.Run(fmt.Sprintf("CRC64_%d", size), func(b *testing.B) {
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
		b.Run(fmt.Sprintf("FNV_%d", size), func(b *testing.B) {
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
		b.Run(fmt.Sprintf("Maphash_%d", size), func(b *testing.B) {
			key := strings.Repeat("a", size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hashfuncbench.Maphash(key)
			}
		})
	}
}
