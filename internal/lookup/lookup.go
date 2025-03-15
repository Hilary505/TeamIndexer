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

/* LookupChunkBySimHash efficiently reads an index file line by line to find a matching SimHash without loading the entire file into memory. */
func LookupChunkBySimHash(indexFile string, SimHash string) (*LookupResult, error) {
	file, err := os.Open(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open index file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // Stream-read line by line
	for scanner.Scan() {
		var chunk indexer.Chunk
		err := json.Unmarshal(scanner.Bytes(), &chunk) // Decode each line
		if err != nil {
			return nil, fmt.Errorf("failed to parse index file: %v", err)
		}

		if SimHash == chunk.Data { // Compare SimHash (modify condition if needed)
			result := &LookupResult{
				SourceFile: chunk.Source,
				Position:   chunk.ID,
				Phrase:     chunk.Data,
			}
			return result, nil // Return result and nil error
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading index file: %v", err)
	}

	return nil, errors.New("SimHash not found in index")
}
