package indexer

import (
	"fmt"
	"hash/fnv"
)

const (
	HashSize = 64
)

type Chunk struct {
	Source string `json:"Source"`
	Data   string `json:"Data"`
	ID     int    `json:"ID"`
}

var ChunkSlice = map[string]*Chunk{}

/* SimHash computes the SimHash fingerprint for a given chunk of data. */
func SimHash(data []byte) string {
	vector := make([]int, HashSize)

	/*Hash the data using FNV-1a */
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
	return fmt.Sprintf("%x", fingerprint)
}
