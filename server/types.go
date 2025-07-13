package server

// Common type definitions that are shared across multiple files

// StoreCipher represents the cipher used for encrypted store data
type StoreCipher int

// msgResults represents the results of message processing
type msgResults struct {
	// Add appropriate fields based on your implementation
	subject   []byte
	reply     []byte
	delivered int
	queues    []string
	// Add any other fields needed
}
