package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"index/internal/chunker"
	"index/internal/indexer"
	"index/internal/lookup"
)

/*===========================================================================
======================== Main Function ======================================
============================================================================== */
func main() {
	currentdir, err := os.Getwd()
	if err != nil {
		log.Println("error: must be working directory")
		return
	}
	command := flag.String("c", "", "Command to execute: 'index' or 'lookup'")
	inputFile := flag.String("i", currentdir + "/large_text.txt", "input file")
	chunkSize := flag.Int("s", 4096, "Size of each chunk in bytes")
	indexFile := flag.String("o", "index.idx", "Path to save or load the index file")
	lookupHash := flag.String("h", "", "SimHash value to lookup")
	flag.Parse()

	if *command != "index" && *command != "lookup" {
		log.Println("Invalid command, use 'index' or 'lookup'")
		return
	} else if *chunkSize%8 != 0 {
		log.Println("Error:Invalid Chunksize")
		return
	}

	// Execute the command
	switch *command {
	case "index":
		if *inputFile == "" {
			log.Println("Check the file path and try again")
			return
		}
		indexCommand(*inputFile, *chunkSize, *indexFile)
	case "lookup":
		if *lookupHash == "" {
			log.Println("SimHash value is required for lookup")
			return
		}
		lookupCommand(*indexFile, *lookupHash)
	}
}

/* indexCommand processes a text file and builds an index */
func indexCommand(inputFile string, chunkSize int, indexFile string) {
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
func lookupCommand(indexFile string, SimHash string) {
	result, err := lookup.LookupChunkBySimHash(indexFile, SimHash)
	if err != nil {
		log.Fatalf("Lookup failed: %v", err)
		return
	}
	fmt.Printf("Source File: %s\n", result.SourceFile)
	fmt.Printf("Position: %d\n", result.Position)
	fmt.Printf("Phrase: %s\n", result.Phrase)
}
