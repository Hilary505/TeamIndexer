package lookup

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"index/internal/indexer"
)

type LookupResult struct {
	SourceFile string `json:"Source"`
	Position   int    `json:"Position"`
	Phrase     string `json:"Phrase"`
}

func LookupChunkBySimHash(indexFile string, SimHash string) (*LookupResult, error) {
	data, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %v", err)
	}
	var chunkMap map[uint64]*indexer.Chunk
	if err := json.Unmarshal(data, &chunkMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal index file: %v", err)
	}

	/* Convert the provided SimHash string to uint64 */
	simHashValue, err := strconv.ParseUint(SimHash, 10, 64) // Use base 10 for decimal input
	if err != nil {
		return nil, fmt.Errorf("invalid SimHash value: %v", err)
	}
	chunk, exists := chunkMap[simHashValue]
	if !exists {
		return nil, errors.New("chunk with the given SimHash not found")
	}
	result := &LookupResult{
		SourceFile: chunk.Source,
		Position:   chunk.ID,
		Phrase:     chunk.Data,
	}
	return result, nil
}