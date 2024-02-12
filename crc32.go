package hashfuncbench

import (
	"hash/crc32"
)

func Crc32(data string) uint32 {
	h := crc32.NewIEEE()
	h.Write([]byte(data))

	return h.Sum32()
}
