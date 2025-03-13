package indexer

import (
	"hash/fnv"
)

const (
	HashSize = 64
)

type Chunk struct {
	Source string
	Data   []byte
	ID     int
	Hash   uint64
}

var (
	ChunkSlice = []*Chunk{}
	ID         = 0
)

// SimHash computes the SimHash fingerprint for a given chunk of data.
func SimHash(data []byte) *Chunk {
	vector := make([]int, HashSize)

	// Hash the data using FNV-1a
	hasher := fnv.New64a()
	hasher.Write(data)
	hash := hasher.Sum64()
	for i := 0; i < HashSize; i++ {
		if (hash>>i)&1 == 1 {
			vector[i]++
		} else {
			vector[i]--
		}
	}
	var fingerprint uint64
	for i := 0; i < HashSize; i++ {
		if vector[i] > 0 {
			fingerprint |= 1 << i
		}
	}
	ChunkSlice = append(ChunkSlice, &Chunk{Source: "", Data: data, ID: ID + 1, Hash: fingerprint})
	return &Chunk{Source: "", Data: data, ID: ID + 1, Hash: fingerprint}
}

// Distance calculates the Hamming distance between two SimHash fingerprints
func Distance(a, b uint64) int {
	xor := a ^ b
	distance := 0
	for xor != 0 {
		distance += int(xor & 1)
		xor >>= 1
	}
	return distance
}
