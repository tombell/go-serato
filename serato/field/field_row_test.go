package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewRowField(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	row, err := field.NewRowField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewRowField err to be nil")
	}

	if row == nil {
		t.Fatal("expected row to not be nil")
	}
}

func TestNewRowFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000010000000400000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewRowField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewRowField err to be io.ErrUnexpectedEOF")
	}
}

func TestNewRowFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000200000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = field.NewRowField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewRowField err to be field.ErrUnexpectedIdentifier")
	}
}

func TestRowValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	row, err := field.NewRowField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewRowField err to be nil")
	}

	actual := row.Value()
	expected := 212

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestRowString(t *testing.T) {
	data, _ := hex.DecodeString("0000000100000004000000D4")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	row, err := field.NewRowField(hdr, buf)
	if err != nil {
		t.Fatal("expected NewRowField err to be nil")
	}

	actual := row.String()
	expected := "212"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
