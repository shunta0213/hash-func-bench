package hashfuncbench_test

import (
	"testing"

	hashfuncbench "github.com/shunta0213/hash-func-bench"
)

var testcase = []string{
	"0",
	"abc",
	"qwer1234",
}

func FuzzSHA256(f *testing.F) {
	for _, data := range testcase {
		f.Add(data)
	}

	seenHashes := make(map[string]string)
	f.Fuzz(func(t *testing.T, data string) {
		hashBytes := hashfuncbench.Sha256(data)
		hash := string(hashBytes)
		if previousData, exists := seenHashes[hash]; exists && previousData != data {
			t.Fatalf("Collision found: data1 = %s, data2 = %s, hash = %s", previousData, data, hash)
		}
		seenHashes[hash] = data
	})
}

func FuzzCRC32(f *testing.F) {
	for _, data := range testcase {
		f.Add(data)
	}

	seenHashes := make(map[uint32]string)
	f.Fuzz(func(t *testing.T, data string) {
		hash := hashfuncbench.Crc32(data)
		if previousData, exists := seenHashes[hash]; exists && previousData != data {
			t.Fatalf("Collision found: data1 = %s, data2 = %s, hash = %d", previousData, data, hash)
		}
		seenHashes[hash] = data
	})
}

func FuzzCRC64(f *testing.F) {
	for _, data := range testcase {
		f.Add(data)
	}

	seenHashes := make(map[uint64]string)
	f.Fuzz(func(t *testing.T, data string) {
		hash := hashfuncbench.Crc64(data)
		if previousData, exists := seenHashes[hash]; exists && previousData != data {
			t.Fatalf("Collision found: data1 = %s, data2 = %s, hash = %d", previousData, data, hash)
		}
		seenHashes[hash] = data
	})
}

func FuzzFNV(f *testing.F) {
	for _, data := range testcase {
		f.Add(data)
	}

	seenHashes := make(map[uint64]string)
	f.Fuzz(func(t *testing.T, data string) {
		hash := hashfuncbench.Fnv(data)
		if previousData, exists := seenHashes[hash]; exists && previousData != data {
			t.Fatalf("Collision found: data1 = %s, data2 = %s, hash = %d", previousData, data, hash)
		}
		seenHashes[hash] = data
	})
}

func FuzzMapHash(f *testing.F) {
	for _, data := range testcase {
		f.Add(data)
	}

	seenHashes := make(map[uint64]string)
	f.Fuzz(func(t *testing.T, data string) {
		hash := hashfuncbench.Maphash(data)
		if previousData, exists := seenHashes[hash]; exists && previousData != data {
			t.Fatalf("Collision found: data1 = %s, data2 = %s, hash = %d", previousData, data, hash)
		}
		seenHashes[hash] = data
	})
}
