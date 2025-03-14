package chunker

import (
	"reflect"
	"testing"
)

func TestNewChunker(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		want    *Chunker
		wantErr bool
	}{
		{"Valid Chunk Size", 5, &Chunker{ChunkSize: 5}, false},
		{"Zero Chunk Size", 0, nil, true},
		{"Negative Chunk Size", -5, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewChunker(tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewChunker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChunker() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mock Indexer Package for Testing
type MockIndexer struct{}

func (m *MockIndexer) SimHash(data []byte) string {
	return "mockHash"
}
func TestChunker_Chunk(t *testing.T) {
	tests := []struct {
		name      string
		chunkSize int
		data      []byte
		want      [][]byte
	}{
		{"Exact Multiples", 3, []byte("abcdef"), [][]byte{[]byte("abc"), []byte("def")}},
		{"Leftover Data", 3, []byte("abcdefg"), [][]byte{[]byte("abc"), []byte("def"), []byte("g")}},
		{"Smaller Than Chunk", 5, []byte("abc"), [][]byte{[]byte("abc")}},
		{"Empty Data", 4, []byte(""), [][]byte{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chunker{ChunkSize: tt.chunkSize}
			c.Chunk(tt.data, "testFile")
		})
	}
}