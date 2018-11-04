package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewDeckField(t *testing.T) {
	data, _ := hex.DecodeString("0000001F0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	deck, err := field.NewDeckField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewDeckField err to be nil, got %v", err)
	}

	if deck == nil {
		t.Fatal("expected deck to not be nil")
	}
}

func TestNewDeckFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001F00000004000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewDeckField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewDeckField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewDeckFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000001D0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewDeckField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewDeckField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestDeckValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001F0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	deck, err := field.NewDeckField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewDeckField err to be nil, got %v", err)
	}

	actual := deck.Value()
	expected := 1

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestDeckString(t *testing.T) {
	data, _ := hex.DecodeString("0000001F0000000400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	deck, err := field.NewDeckField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewDeckField err to be nil, got %v", err)
	}

	actual := deck.String()
	expected := "1"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
