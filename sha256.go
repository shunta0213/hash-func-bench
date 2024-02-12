package hashfuncbench

import (
	"crypto/sha256"
)

func Sha256(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))

	return h.Sum(nil)
}
