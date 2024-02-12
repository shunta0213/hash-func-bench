package hashfuncbench

import (
	"hash/crc64"
)

func Crc64(data string) uint64 {
	h := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(data), h)
}
