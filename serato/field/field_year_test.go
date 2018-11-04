package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewYearField(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	year, err := field.NewYearField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewYearField err to be nil, got %v", err)
	}

	if year == nil {
		t.Fatal("expected year to not be nil")
	}
}

func TestNewYearFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A0032003000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewYearField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewYearField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewYearFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000001E0000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewYearField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewYearField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestYearValue(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	year, err := field.NewYearField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewYearField err to be nil, got %v", err)
	}

	actual := year.Value()
	expected := "2018"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestYearString(t *testing.T) {
	data, _ := hex.DecodeString("000000170000000A00320030003100380000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	year, err := field.NewYearField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewYearField err to be nil, got %v", err)
	}

	actual := year.String()
	expected := "2018"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
