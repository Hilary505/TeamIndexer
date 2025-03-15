package indexer

import (
	"fmt"
	"hash/fnv"
	"sync"
)

const (
	HashSize = 64
)

var ChunkSlice = map[string]*Chunk{}
type Chunk struct {
	Source string `json:"Source"`
	Data   string `json:"Data"`
	ID     int    `json:"ID"`
}

// SimHash computes the SimHash fingerprint for a given chunk of data efficiently.
func SimHash(data []byte) string {
	vector := make([]int, HashSize) // Bit vector
	var mu sync.Mutex               // Mutex for safe concurrent writes
	var wg sync.WaitGroup            // WaitGroup to manage goroutines

	// Divide data into smaller parts for parallel hashing
	chunkSize := len(data) / 4 // Divide into 4 parts (adjust for larger datasets)
	if chunkSize < 1 {
		chunkSize = len(data) // Process as a single chunk if small
	}

	for i := 0; i < len(data); i += chunkSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + chunkSize
			if end > len(data) {
				end = len(data)
			}

			// Hash the chunk
			hasher := fnv.New64a()
			hasher.Write(data[start:end])
			hash := hasher.Sum64()

			// Update the bit vector
			mu.Lock()
			for i := 0; i < HashSize; i++ {
				if (hash>>i)&1 == 1 {
					vector[i]++
				} else {
					vector[i]--
				}
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait() // Ensure all goroutines complete before proceeding

	// Generate fingerprint
	var fingerprint uint64
	for i := 0; i < HashSize; i++ {
		if vector[i] > 0 {
			fingerprint |= 1 << i
		}
	}
	return fmt.Sprintf("%x", fingerprint)
}
