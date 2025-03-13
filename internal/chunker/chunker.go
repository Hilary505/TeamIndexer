package chunker

import(
	"errors"
)

type Chunker struct {
	chunkSize int
}

func NewChunker(chunkSize int) (*Chunker, error){
	if chunkSize <= 0 {
		return nil, errors.New("chunk size must be greater than 0")
	}
	return &Chunker{chunkSize: chunkSize},nil
}

func (c *Chunker) Chunk(data []byte) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i+=c.chunkSize {
		end := i + c.chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks,data[i:end])
	}
	return chunks
}