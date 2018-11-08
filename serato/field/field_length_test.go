package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

// XXX: Length field appears to always be empty in session files.

func TestNewLengthField(t *testing.T) {
	data, _ := hex.DecodeString("0000000A0000001200300030003A00310038003A003000310000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	length, err := field.NewLengthField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLengthField err to be nil, got %v", err)
	}

	if length == nil {
		t.Fatal("expected length to not be nil")
	}
}

func TestNewLengthFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000A0000001200300030003A00310038003A0030003100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLengthField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewLengthField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewLengthFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000B0000001200300030003A00310038003A003000310000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLengthField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewLengthField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestLengthValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000A0000001200300030003A00310038003A003000310000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	length, err := field.NewLengthField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLengthField err to be nil, got %v", err)
	}

	actual := length.Value()
	expected := "00:18:01"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestLengthString(t *testing.T) {
	data, _ := hex.DecodeString("0000000A0000001200300030003A00310038003A003000310000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	length, err := field.NewLengthField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLengthField err to be nil, got %v", err)
	}

	actual := length.String()
	expected := "00:18:01"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
