package indexer

import (
	"hash/fnv"
)

const(
	HashSize = 64
)

// SimHash computes the SimHash fingerprint for a given chunk of data.
func SimHash(data  []byte) uint64 {
	vector := make([]int, HashSize)

	//Hash the data using FNV-1a
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
	return fingerprint
}

//Distance calculates the Hamming distance between two SimHash fingerprints
func Distance(a, b uint64) int {
	xor := a ^ b
	distance := 0
	for xor != 0 {
		distance += int(xor & 1)
		xor >>= 1
	}
	return distance
}