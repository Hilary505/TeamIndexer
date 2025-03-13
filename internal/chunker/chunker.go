package chunker

import(
	"errors"
)

type Chunker struct {
	ChunkSize int
}

func NewChunker(ChunkSize int) (*Chunker, error){
	if ChunkSize <= 0 {
		return nil, errors.New("chunk size must be greater than 0")
	}
	return &Chunker{ChunkSize: ChunkSize},nil
}

func (c *Chunker) Chunk(data []byte) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += c.ChunkSize {
		end := i + c.ChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks,data[i:end])
	}
	return chunks
}