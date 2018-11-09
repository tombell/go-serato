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

// This will test all fields that Serato DJ Pro 2.0.5 will include values for.
func TestNewFieldsKnownFields(t *testing.T) {
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

// This will test all fields that Serato DJ Pro 2.0.5 will leave empty, and not
// include in the history file.
func TestNewFieldsWithKnownEmptyFields(t *testing.T) {
	t.Skip()
}

// This will test all fields that are parsed, but we don't know what they are
// used for.
func TestNewFieldsWithHandledUnkownFields(t *testing.T) {
	t.Skip()
}

// This will test that we skip passed any unknown fields which are not
// specifically handled in code.
func TestNewFieldsSkipsUnhandledUnknownFields(t *testing.T) {
	t.Skip()
}

func TestNewFieldsUnexpectedEOF(t *testing.T) {
	t.Skip()
}

// TODO: add tests, maybe table test that NewXXXField errors in NewFields
