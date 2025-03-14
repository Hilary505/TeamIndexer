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
	"index/internal/utils"

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
		utils.IndexCommand(*inputFile, *chunkSize, *indexFile)
	case "lookup":
		if *lookupHash == "" {
			log.Println("SimHash value is required for lookup")
			return
		}
		utils.LookupCommand(*indexFile, *lookupHash)
	}
}

