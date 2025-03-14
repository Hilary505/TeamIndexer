package chunker

import (
	"errors"
	"fmt"

	"index/internal/indexer"
)

type Chunker struct {
	ChunkSize int
}

func NewChunker(ChunkSize int) (*Chunker, error) {
	if ChunkSize <= 0 {
		return nil, errors.New("chunk size must be greater than 0")
	}
	return &Chunker{ChunkSize: ChunkSize}, nil
}

func (c *Chunker) Chunk(data []byte, inputFile string) {
	var chunks [][]byte

	for i := 0; i < len(data); i += c.ChunkSize {
		end := i + c.ChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	// return chunks
	for j, chunk := range chunks {
		simhash := indexer.SimHash(chunk)
		fmt.Println(inputFile)
		indexer.ChunkSlice[simhash] = &indexer.Chunk{Source: inputFile, Data: string(chunk), ID: j + 1}
	}
}
