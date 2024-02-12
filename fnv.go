package hashfuncbench

import (
	"hash/fnv"
)

func Fnv(data string) uint64 {
	h := fnv.New64()
	h.Write([]byte(data))

	return h.Sum64()
}
