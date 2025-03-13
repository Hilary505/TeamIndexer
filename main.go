package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	//"index/internal/chunker"
	"index/internal/indexer"
)

func main() {
	currentdir, _ := os.Getwd()
	command := flag.String("c", "", "Command to execute: 'index' or 'lookup'")
	inputFile := flag.String("i", currentdir+"/internal/testdata/large_text.txt", "~/media/enungo/ed/TeamIndexer/internal/testdata/large_text.txt")
	chunkSize := flag.Int("s", 4096, "Size of each chunk in bytes")
	indexFile := flag.String("o", "index.idx", "Path to save or load the index file")
	lookupHash := flag.String("hash", "", "SimHash value to lookup")
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

	indexer.SimHash(data)
	json.Unmarshal(data, &indexer.ChunkSlice)
	// chunker, err := chunker.NewChunker((chunkSize))
	// if err != nil {
	// 	log.Printf("Failed to create chunker: %v", err)
	// 	os.Exit(1)
	// }
	// idx := indexer.NewIndexer(chunker)
	// fingerprints := idx.Process(data)
	// // Build the index
	// index := idx.BuildIndex(data, fingerprints)
	file, err := os.Create(indexFile)
	if err != nil {
		log.Printf("Failed to create index file: %v", err)
	}
	defer file.Close()

	// encoder := gob.NewEncoder(file)
	// if err := encoder.Encode(index); err != nil {
	// 	log.Printf("Failed to encode index: %v", err)
	// 	os.Exit(1)
	// }
	jsonw, _ := json.Marshal(indexer.ChunkSlice)
	os.WriteFile(indexFile, jsonw, 0o0644)
	fmt.Printf("Index saved to %s\n", indexFile)
}
