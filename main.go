package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"index/internal/chunker"
	"index/internal/indexer"
)

func main() {
	currentdir, _ := os.Getwd()
	command := flag.String("c", "", "Command to execute: 'index' or 'lookup'")
	inputFile := flag.String("i", currentdir+"/internal/testdata/large_text.txt", "")
	chunkSize := flag.Int("s", 4096, "Size of each chunk in bytes")
	indexFile := flag.String("o", "index.idx", "Path to save or load the index file")
	lookupHash := flag.String("h", "", "SimHash value to lookup")
	flag.Parse()

	if *command != "index" && *command != "lookup" {
		log.Println("Invalid command, use 'index' or 'lookup'")
	}

	// Execute the command
	switch *command {
	case "index":
		if *inputFile == "" {
			log.Println("Check the file path and try again")
		}
		indexCommand(*inputFile, *chunkSize, *indexFile)
	case "lookup":
		if *lookupHash == "" {
			log.Println("SimHash value is required for lookup")
		}
		// lookupCommand(*indexFile, *lookupHash)
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
	chunker.Chunk(data)
	json.Unmarshal(data, &indexer.ChunkSlice)

	file, err := os.Create(indexFile)
	if err != nil {
		log.Printf("Failed to create index file: %v", err)
	}
	defer file.Close()
	jsonw, _ := json.Marshal(indexer.ChunkSlice)
	os.WriteFile(indexFile, jsonw, 0o0644)
	fmt.Printf("Index saved to %s\n", indexFile)
	for _, each := range indexer.ChunkSlice {
		fmt.Println(each.ID)
	}
}
