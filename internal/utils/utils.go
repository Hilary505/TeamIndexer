package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"index/internal/chunker"
	"index/internal/indexer"
	"index/internal/lookup"
)

/* indexCommand processes a text file and builds an index */
func IndexCommand(inputFile string, chunkSize int, indexFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Printf("Failed to read input file: %v", err)
		os.Exit(1)
	}
	chunker, err := chunker.NewChunker(chunkSize)
	if err != nil {
		log.Printf("Failed to create chunker: %v", err)
		os.Exit(1)
	}
	chunker.Chunk(data, inputFile)
	json.Unmarshal(data, &indexer.ChunkSlice)
	file, err := os.Create(indexFile)
	if err != nil {
		log.Printf("Failed to create index file: %v", err)
		return
	}
	defer file.Close()
	jsonw, err := json.Marshal(indexer.ChunkSlice)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	os.WriteFile(indexFile, jsonw, 0o0644)
	fmt.Printf("Index saved to %s\n", indexFile)
	for simHash, chunk := range indexer.ChunkSlice {
		fmt.Printf("Chunk ID: %d, SimHash:%v\n", chunk.ID, simHash)
	}
}

/* lookupCommand looks up a chunk by its SimHash value */
func LookupCommand(indexFile string, SimHash string) {
	result, err := lookup.LookupChunkBySimHash(indexFile, SimHash)
	if err != nil {
		log.Fatalf("Lookup failed: %v", err)
		return
	}
	fmt.Printf("Source File: %s\n", result.SourceFile)
	fmt.Printf("Position: %d\n", result.Position)
	fmt.Printf("Phrase: %s\n", result.Phrase)
}
