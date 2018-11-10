package field_test

import (
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

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

	if fields.Row == nil {
		t.Fatal("expected fields.Row to not be nil")
	}

	if fields.FullPath == nil {
		t.Fatal("expected fields.FullPath to not be nil")
	}

	if fields.Location != nil {
		t.Fatalf("expected fields.Location to be nil, got %v", fields.Location)
	}

	if fields.Filename != nil {
		t.Fatalf("expected fields.Filename to be nil, got %v", fields.Filename)
	}

	if fields.Title == nil {
		t.Fatal("expected fields.Title to not be nil")
	}

	if fields.Artist == nil {
		t.Fatal("expected fields.Artist to not be nil")
	}

	if fields.Album == nil {
		t.Fatal("expected fields.Album to not be nil")
	}

	if fields.Genre == nil {
		t.Fatal("expected fields.Genre to not be nil")
	}

	if fields.Length != nil {
		t.Fatalf("expected fields.Length to be nil, got %v", fields.Length)
	}

	if fields.Bitrate != nil {
		t.Fatalf("expected fields.Bitrate to be nil, got %v", fields.Bitrate)
	}

	if fields.BPM == nil {
		t.Fatal("expected fields.BPM to not be nil")
	}

	if fields.Comment == nil {
		t.Fatal("expected fields.Comment to not be nil")
	}

	if fields.Grouping == nil {
		t.Fatal("expected fields.Grouping to not be nil")
	}

	if fields.Remixer == nil {
		t.Fatal("expected fields.Remixer to not be nil")
	}

	if fields.Label == nil {
		t.Fatal("expected fields.Label to not be nil")
	}

	if fields.Comment == nil {
		t.Fatal("expected fields.Comment to not be nil")
	}

	if fields.Year == nil {
		t.Fatal("expected fields.Year to not be nil")
	}

	if fields.StartTime == nil {
		t.Fatal("expected fields.StartTime to not be nil")
	}

	if fields.EndTime == nil {
		t.Fatal("expected fields.EndTime to not be nil")
	}

	if fields.Deck == nil {
		t.Fatal("expected fields.Deck to not be nil")
	}

	if fields.PlayTime == nil {
		t.Fatal("expected fields.PlayTime to not be nil")
	}

	if fields.SessionID == nil {
		t.Fatal("expected fields.SessionID to not be nil")
	}

	if fields.Played == nil {
		t.Fatal("expected fields.Played to not be nil")
	}

	if fields.Key == nil {
		t.Fatal("expected fields.Key to not be nil")
	}

	if fields.Added == nil {
		t.Fatal("expected fields.Added to not be nil")
	}

	if fields.UpdatedAt == nil {
		t.Fatal("expected fields.UpdatedAt to not be nil")
	}
}

// This will test all fields that Serato DJ Pro 2.0.5 will leave empty, and not
// include in the history file. Use handwritten data to test those fields.
func TestNewFieldsWithKnownEmptyFields(t *testing.T) {
	str := readTestDataFile(t, "./testdata/known-empty-fields.txt")
	data, _ := hex.DecodeString(str)

	fields, err := field.NewFields(data)
	if err != nil {
		t.Fatalf("expected NewFields err to be nil, got %v", err)
	}

	if fields == nil {
		t.Fatal("expected fields to not be nil")
	}

	if fields.Location == nil {
		t.Fatal("expected fields.Location to not be nil")
	}

	if fields.Filename == nil {
		t.Fatal("expected fields.Filename to not be nil")
	}

	if fields.Length == nil {
		t.Fatal("expected fields.Length to not be nil")
	}

	if fields.Bitrate == nil {
		t.Fatal("expected fields.Bitrate to not be nil")
	}
}

// This will test that we skip passed any unknown fields which are not
// specifically handled in code.
func TestNewFieldsSkipsUnhandledUnknownFields(t *testing.T) {
	str := readTestDataFile(t, "./testdata/unknown-unhandled-fields.txt")
	data, _ := hex.DecodeString(str)

	fields, err := field.NewFields(data)
	if err != nil {
		t.Fatalf("expected NewFields err to be nil, got %v", err)
	}

	if fields == nil {
		t.Fatal("expected fields to not be nil")
	}

	if fields.Row == nil {
		t.Fatal("expected fields.Row to not be nil")
	}
}

func TestNewFieldsUnexpectedEOF(t *testing.T) {
	str := readTestDataFile(t, "./testdata/new-fields-error.txt")
	data, _ := hex.DecodeString(str)

	_, err := field.NewFields(data)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewFields err to be io.UnexpectedEOF, got %v", err)
	}
}

// TODO: add tests, maybe table test that NewXXXField errors in NewFields
