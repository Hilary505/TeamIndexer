package lookup

import (
	"bufio"
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

/* LookupChunkBySimHash reads an index file line by line, unmarshals each line, and retrieves a chunk based on a given SimHash. */
func LookupChunkBySimHash(indexFile string, SimHash string) (*LookupResult, error) {
	file, err := os.Open(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open index file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var chunkMap map[string]*indexer.Chunk
		if err := json.Unmarshal(scanner.Bytes(), &chunkMap); err != nil {
			return nil, fmt.Errorf("failed to parse index file: %v", err)
		}

		if chunk, exists := chunkMap[SimHash]; exists {
			return &LookupResult{
				SourceFile: chunk.Source,
				Position:   chunk.ID,
				Phrase:     chunk.Data,
			}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading index file: %v", err)
	}

	return nil, errors.New("SimHash not found in index")
}
