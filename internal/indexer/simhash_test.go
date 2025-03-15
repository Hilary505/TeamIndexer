package indexer

import (
	"testing"
)

// TestSimHashConsistency ensures the SimHash function returns the same result for the same input
func TestSimHashConsistency(t *testing.T) {
	data := []byte("This is a test string")
	hash1 := SimHash(data)
	hash2 := SimHash(data)

	if hash1 != hash2 {
		t.Errorf("Expected SimHash to be consistent, but got different hashes: %s and %s", hash1, hash2)
	}
}

// TestSimHashDifferentInputs ensures different inputs produce different hashes
func TestSimHashDifferentInputs(t *testing.T) {
	data1 := []byte("This is a test string")
	data2 := []byte("This is another test string")

	hash1 := SimHash(data1)
	hash2 := SimHash(data2)

	if hash1 == hash2 {
		t.Errorf("Expected different inputs to produce different SimHashes, but got the same: %s", hash1)
	}
}

// TestSimHashParallelExecution checks if SimHash runs correctly in parallel
func TestSimHashParallelExecution(t *testing.T) {
	data := []byte("Parallel execution test string")
	expectedHash := SimHash(data) // Get expected output in a single run

	t.Run("ParallelTest", func(t *testing.T) {
		t.Parallel()
		hash := SimHash(data)

		if hash != expectedHash {
			t.Errorf("Expected SimHash %s, but got %s", expectedHash, hash)
		}
	})
}

// BenchmarkSimHash tests the performance of SimHash with large data
func BenchmarkSimHash(b *testing.B) {
	data := []byte("Benchmark test data for SimHash function")
	for i := 0; i < b.N; i++ {
		SimHash(data)
	}
}
