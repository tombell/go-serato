package field_test

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func generateBytes(t *testing.T, data string) []byte {
	t.Helper()

	b, err := hex.DecodeString(data)
	if err != nil {
		t.Fatalf("failed to generate bytes from: %v (%v)", data, err)
	}

	return b
}

func readTestDataFile(t *testing.T, filepath string) string {
	t.Helper()

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("failed to read bytes from: %v (%v)", filepath, err)
	}

	return string(data)
}
