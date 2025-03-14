package lookup

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"index/internal/indexer"
)

type LookupResult struct {
	SourceFile string `json:"Source"`
	Position   int    `json:"Position"`
	Phrase     string `json:"Phrase"`
}

/*function to look for chunk based on the SimHash */
func LookupChunkBySimHash(indexFile string, SimHash string) (*LookupResult, error) {
	data, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %v", err)
	}
	var chunkMap map[string]*indexer.Chunk
	if err := json.Unmarshal(data, &chunkMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal index file: %v", err)
	}

	/* Convert the provided SimHash string to uint64 */
	simHashValue := SimHash
	chunk, exists := chunkMap[simHashValue]
	if !exists {
		return nil, errors.New("Error:SimHash not found")
	}
	result := &LookupResult{
		SourceFile: chunk.Source,
		Position:   chunk.ID,
		Phrase:     chunk.Data,
	}
	return result, nil
}