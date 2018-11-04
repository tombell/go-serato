package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewPlayedField(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayedField err to be nil, got %v", err)
	}

	if played == nil {
		t.Fatal("expected played to not be nil")
	}
}

func TestNewPlayedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003200000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewPlayedField(hdr, buf)
	if err != io.EOF {
		t.Fatalf("expected NewPlayedField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewPlayedFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000330000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewPlayedField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewPlayedField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestPlayedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayedField err to be nil, got %v", err)
	}

	actual := played.Value()
	expected := true

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestPlayedString(t *testing.T) {
	data, _ := hex.DecodeString("000000320000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	played, err := field.NewPlayedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayedField err to be nil, got %v", err)
	}

	actual := played.String()
	expected := "true"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
