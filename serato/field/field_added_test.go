package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewAddedField(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAddedField err to be nil, got %v", err)
	}

	if added == nil {
		t.Fatal("expected added to not be nil")
	}
}

func TestNewAddedFieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003400000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewAddedField(hdr, buf)
	if err != io.EOF {
		t.Fatalf("expected NewAddedField err to be io.EOF, got %v", err)
	}
}

func TestNewAddedFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000330000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewAddedField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewAddedField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestAddedValue(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAddedField err to be nil, got %v", err)
	}

	actual := added.Value()
	expected := false

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestAddedString(t *testing.T) {
	data, _ := hex.DecodeString("000000340000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	added, err := field.NewAddedField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAddedField err to be nil, got %v", err)
	}

	actual := added.String()
	expected := "false"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
