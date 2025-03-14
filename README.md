# **Jam: Fast & Scalable Text Indexer**

## **Overview**
Efficiently searching and retrieving data from large text files is a common challenge in computing. This project implements a **fast and scalable file indexing system** in Go that can:
1. Parse a text file and split it into fixed-size chunks (e.g., 1KB, 4KB, etc.).
2. Generate a **SimHash fingerprint** for each chunk to group similar chunks together.
3. Build an in-memory index that maps SimHash values to byte offsets for fast retrieval.
4. Allow quick lookups based on chunk hashes.

The implementation prioritizes **speed**, **efficient memory use**, and **clean code structure**.

## **Features**

1. **Indexing**:
   - Splits large text files into fixed-size chunks.
   - Computes SimHash fingerprints for each chunk.
   - Creates an in-memory index mapping SimHash values to byte offsets.
   - Saves the index to a file for persistent storage.

2. **Lookup**:
   - Retrieves the position of a chunk in the original file based on its SimHash value.
   - Outputs:
     - The original source file.
     - The byte offset of the chunk in the file.
     - The content (or phrase) associated with the chunk.

3. **Error Handling**:
   - Handles missing files, invalid chunk sizes, or missing SimHash values gracefully.
   - Provides meaningful error messages to guide users.

4. **Extensibility**:
   - Modular design makes it easy to extend or modify components (e.g., add multi-threading or fuzzy search).

## **Project Structure**
```plaintext
textindex/
├── cmd/
│   └── main.go         # CLI entry point
├── internal/
│   ├── chunker/
│   │   └── chunker.go  # Chunking logic
│   ├── indexer/
│   │   ├── indexer.go  # Indexing logic
│   │   └── simhash.go  # SimHash implementation
│   └── lookup/
│       └── lookup.go   # Lookup functionality
├── testdata/
│   └── large_text.txt  # Sample text file for testing
└── go.mod              # Go module file
```

---

## **Installation**

1. Clone the repository:
 ```bash
    git clone https://github.com/Hilary505/TeamIndexer.git

    cd TeamIndexer
```

2. Install dependencies:
 ```bash
    go mod tidy
 ```

3. Build the executable:
```bash
    go build -o textindex ./cmd/main.go
```

## **Usage**

### **Indexing a Text File**
The `index` command processes a text file, splits it into chunks, computes SimHash fingerprints, and creates an index.

#### Syntax:
```bash
textindex index -i  -s  -o 
```

#### Arguments:
| Argument         | Description                                                      |
|------------------|------------------------------------------------------------------|
| `-i`             | Path to the input text file.                                     |
| `-s`             | Size of each chunk in bytes (default: 4096 bytes).               |
| `-o`             | Path to save the generated index file (default: `index.idx`).    |

#### Example Usage:
```bash
textindex index -i testdata/large_text.txt -s 4096 -o index.idx
```
This command splits `large_text.txt` into 4KB chunks, generates SimHash fingerprints, and saves the index in `index.idx`.

### **Looking Up a Chunk by SimHash**
The `lookup` command retrieves information about a chunk based on its SimHash value.

#### Syntax:
```bash
textindex lookup -i  -h 
```

#### Arguments:
| Argument         | Description                                                      |
|------------------|------------------------------------------------------------------|
| `-i`             | Path to the previously generated index file.                     |
| `-h`             | The SimHash value of the chunk to search for.                    |

#### Example Usage:
```bash
textindex lookup -i index.idx -h 3e4f1b2c98a6...
```
This command finds the position of the chunk with the given SimHash and returns its byte offset in the original file along with its content.

## **Implementation Details**

### **1. Chunking Logic (`chunker.go`)**
The `Chunker` splits input data into fixed-size chunks using a configurable chunk size.

#### Key Features:
- Handles edge cases like empty input or partial final chunks.
- Returns slices of byte arrays representing individual chunks.

### **2. SimHash Fingerprinting (`simhash.go`)**
The `SimHash` function generates a unique fingerprint for each chunk based on its content.

#### Key Features:
- Uses FNV-1a hash function for efficient hashing.
- Aggregates bit vectors to compute a compact fingerprint.
- Ensures similar chunks produce similar hash values.

### **3. Indexing Logic (`indexer.go`)**
The `Indexer` processes data, computes fingerprints, and builds an in-memory index.

#### Key Features:
- Maps SimHash values to byte offsets for fast lookup.
- Supports efficient serialization and deserialization using Go's `encoding/gob`.

### **4. Lookup Logic (`lookup.go`)**
The lookup functionality retrieves information about a specific chunk based on its SimHash value.

#### Outputs:
- Original source file name.
- Byte offset of the matching chunk.
- Content of the matching chunk.

## **Error Handling**

### Common Errors and Fixes:

| Error                  | Cause                             | Suggested Fix                                      |
|------------------------|-----------------------------------|--------------------------------------------------|
| File not found         | Input file does not exist         | Check the file path and try again                |
| Invalid chunk size     | Chunk size is missing or invalid  | Provide a valid chunk size (e.g., 1024)          |
| SimHash not found      | Hash is not present in the index  | Ensure the file was indexed before looking up    |

---

## **Testing**

Unit tests are included for all major components (`chunker`, `simhash`, `indexer`). To run tests:

```bash
go test ./...
```

## **Authors**
[Hilary]()
[Nungo]()
[Namayi]()