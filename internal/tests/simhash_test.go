package tests

import (
	"testing"
	"index/internal/indexer"
)

func TestSimHash(t *testing.T) {
	tests := []struct {
		data     string
		expected string 
	}{
		{"hello world", "76083cd157048fbc"},
		{"Go is great!", "64df5ffb79241ccb"},
		{"Testing SimHash", "f57d7eadfbf20d53"},
		{"", ""},
	}

	for _, tt := range tests {
		got := indexer.SimHash([]byte(tt.data))
		if len(got) == 0 {
			t.Errorf("SimHash(%q) returned an empty hash", tt.data)
		}
	}
}
