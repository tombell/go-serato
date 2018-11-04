package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewField70Field(t *testing.T) {
	data, _ := hex.DecodeString("000000460000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField70Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField70Field err to be nil, got %v", err)
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField70FieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000004600000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField70Field(hdr, buf)
	if err != io.EOF {
		t.Fatalf("expected NewField70Field err to be io.EOF, got %v", err)
	}
}

func TestNewField70FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000440000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField70Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewField70Field err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestField70Value(t *testing.T) {
	data, _ := hex.DecodeString("000000460000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField70Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField70Field err to be nil, got %v", err)
	}

	actual := field.Value()
	expected := byte(0)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestField70String(t *testing.T) {
	data, _ := hex.DecodeString("000000460000000100")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField70Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField70Field err to be nil, got %v", err)
	}

	actual := field.String()
	expected := "0"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
