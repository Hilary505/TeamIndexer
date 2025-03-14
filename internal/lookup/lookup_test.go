package lookup

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"index/internal/indexer"
)

/* Mock index file content */
func createTempIndexFile(t *testing.T, content map[string]*indexer.Chunk) string {
	data, err := json.Marshal(content)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}
	tmpFile, err := os.CreateTemp("", "index-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := tmpFile.Write(data); err != nil {
		t.Fatalf("failed to write test data to temp file: %v", err)
	}
	tmpFile.Close()
	return tmpFile.Name()
}

func TestLookupChunkBySimHash(t *testing.T) {
	// Sample chunk data
	testChunks := map[string]*indexer.Chunk{
		"1234567890abcdef": {Source: "file1.txt", ID: 1, Data: "sample phrase"},
		"abcdef1234567890": {Source: "file2.txt", ID: 2, Data: "another phrase"},
	}

	// Create a temporary index file
	indexFile := createTempIndexFile(t, testChunks)
	defer os.Remove(indexFile)

	tests := []struct {
		name    string
		index   string
		SimHash string
		want    *LookupResult
		wantErr bool
	}{
		{
			name:    "Valid SimHash",
			index:   indexFile,
			SimHash: "1234567890abcdef",
			want:    &LookupResult{SourceFile: "file1.txt", Position: 1, Phrase: "sample phrase"},
			wantErr: false,
		},
		{
			name:    "Non-existent SimHash",
			index:   indexFile,
			SimHash: "nonexistenthash",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Missing index file",
			index:   "non_existent_file.json",
			SimHash: "1234567890abcdef",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LookupChunkBySimHash(tt.index, tt.SimHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupChunkBySimHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupChunkBySimHash() = %v, want %v", got, tt.want)
			}
		})
	}
}