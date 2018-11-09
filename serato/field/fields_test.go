package field_test

import (
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func readTestDataFile(t *testing.T, filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("failed to read bytes from: %v (%v)", filepath, err)
	}

	return string(data)
}

func TestNewFields(t *testing.T) {
	str := readTestDataFile(t, "./testdata/known-fields.txt")
	data, _ := hex.DecodeString(str)

	fields, err := field.NewFields(data)
	if err != nil {
		t.Fatalf("expected NewFields err to be nil, got %v", err)
	}

	if fields == nil {
		t.Fatal("expected fields to not be nil")
	}
}

func TestNewFieldsWithUnusedFields(t *testing.T) {
	t.Skip()
}

func TestNewFieldsWithUnhandledFields(t *testing.T) {
	t.Skip()
}

func TestNewFieldsUnexpectedEOF(t *testing.T) {
	t.Skip()
}
