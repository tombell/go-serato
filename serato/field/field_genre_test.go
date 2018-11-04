package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewGenreField(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	genre, err := field.NewGenreField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGenreField err to be nil, got %v", err)
	}

	if genre == nil {
		t.Fatal("expected genre to not be nil")
	}
}

func TestNewGenreFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewGenreField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewGenreField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewGenreFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewGenreField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewGenreField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestGenreValue(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	genre, err := field.NewGenreField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGenreField err to be nil, got %v", err)
	}

	actual := genre.Value()
	expected := "House"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}

func TestGenreString(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0048006F0075007300650000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	genre, err := field.NewGenreField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGenreField err to be nil, got %v", err)
	}

	actual := genre.String()
	expected := "House"

	if actual != expected {
		t.Fatalf("expected value to be %s, got %s", expected, actual)
	}
}
