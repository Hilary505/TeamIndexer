package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"index/internal/indexer"
)

func main() {
	indexer.SimHash([]byte("hallo world"))
	jsonsata, _ := json.Marshal(indexer.ChunkSlice)
	os.WriteFile("./textindex.txt", []byte(jsonsata), 0o0644)
	filebyte, _ := os.ReadFile("textindex.txt")
	json.Unmarshal(filebyte, &indexer.ChunkSlice)
	u, _ := strconv.ParseUint("16570556702914959827", 10, 64)
	fmt.Println(*&indexer.ChunkSlice[u].Data)
}
