package chunk_test

import (
	"encoding/hex"
	"testing"
)

// TODO: do we need to test the constants and/or error?

func generateBytes(t *testing.T, data string) []byte {
	t.Helper()

	b, err := hex.DecodeString(data)
	if err != nil {
		t.Fatalf("failed to generate bytes from: %v (%v)", data, err)
	}

	return b
}
