package tests

import (
	"testing"

	"index/internal/indexer"
)

func TestSimHash(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []byte
		expected uint64
	}{
		{
			name:     "Empty text",
			input:    []byte(""),
			expected: 0xcbf29ce484222325,
		},
		{
			name:     "Small text",
			input:    []byte("hello"),
			expected: 0xa430d84680aabd0b,
		},
		{
			name:     "Large text",
			input:    []byte("This is a larger chunk of text for testing SimHash."),
			expected: 0x75198cfdfed76b94,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fingerprint := indexer.SimHash(tt.input)
			if fingerprint != tt.expected {
				t.Errorf("SimHash(%q) = %x; expected %x", tt.input, fingerprint, tt.expected)
			}
		})
	}
}
