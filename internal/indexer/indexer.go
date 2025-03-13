package indexer

import (
	"index/internal/chunker"
)

// Indexer processes data and generates SimHash fingerprints for each chunk
type Indexer struct {
	chunker *chunker.Chunker
}

//NewIndexer creates a new Indexer
func NewIndexer(chunker *chunker.Chunker) *Indexer {
	return &Indexer{chunker: chunker}
}

func (i *Indexer) Process(data []byte) []uint64 {
	chunks := i.chunker.Chunk(data)
	fingerprints := make([]uint64, len(chunks))
	for idx, chunk := range chunks {
		fingerprints[idx] = SimHash(chunk)
	}
	return fingerprints
}

// BuildIndex builds an in-memory index mapping SimHash values to byte offsets.
func (i *Indexer) BuildIndex(data []byte, fingerprints []uint64) map[uint64][]int {
	index := make(map[uint64][]int)
	chunks := i.chunker.Chunk(data)

	for idx, fp := range fingerprints {
		index[fp] = append(index[fp], idx*i.chunker.chunkSize) 
	}
	return index
}