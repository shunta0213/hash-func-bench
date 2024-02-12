package hashfuncbench_test

import (
	"testing"

	hashfuncbench "github.com/shunta0213/hash-func-bench"
)

func BenchmarkSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashfuncbench.Sha256("test")
	}
}

func BenchmarkCRC32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashfuncbench.Crc32("test")
	}
}

func BenchmarkCRC64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashfuncbench.Crc64("test")
	}
}

func BenchmarkFNV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashfuncbench.Fnv("test")
	}
}

func BenchmarkMapHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashfuncbench.Maphash("test")
	}
}
