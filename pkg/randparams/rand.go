package randparams

import (
	"hash/fnv"
	"math/rand"
)

type Rand struct {
	rand.Rand
}

func New(seed []byte) *Rand {

	// hash seed to create a seeded rng
	h := fnv.New64a()
	h.Write(seed)
	r := rand.New(rand.NewSource(int64(h.Sum64())))

	return &Rand{
		Rand: *r,
	}
}

func (r *Rand) Int32InRange(min int32, max int32) int32 {
	return int32(r.Float32InRange(float32(min), float32(max)))
}

func (r *Rand) Float32InRange(min float32, max float32) float32 {
	n := r.Float32() * (max - min)
	return n + min
}
