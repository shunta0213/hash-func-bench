package hashfuncbench

import "hash/maphash"

func Maphash(data string) uint64 {
	h := maphash.Hash{}
	h.WriteString(data)

	return h.Sum64()
}
