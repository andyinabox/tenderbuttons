package params

import (
	"hash/fnv"
	"math/rand"
)

func GetSeededRandom(b []byte) *rand.Rand {
	h := fnv.New64a()
	h.Write(b)
	return rand.New(rand.NewSource(int64(h.Sum64())))
}
