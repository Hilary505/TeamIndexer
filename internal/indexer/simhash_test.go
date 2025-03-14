package indexer

import (
	"testing"
)

func TestSimHash(t *testing.T) {
	tests := []struct {
		name    string
		data     string
		expected string 
	}{
		{"string1","hello world", "779a65e7023cd2e7"},
		{"string2","Go is great!", "8d937047e4b4d2f3"},
		{"string3","Testing SimHash", "3aaa18b79c15c450"},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimHash([]byte(tt.data))
			if got != tt.expected {
				t.Errorf("SimHash(%q) = %q; expected %q", tt.data, got, tt.expected)
			}
		})
	}
}
