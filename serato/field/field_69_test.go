package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewField69Field(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField69Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField69Field err to be nil, got %v", err)
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField69FieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField69Field(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewField69Field err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewField69FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000440000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField69Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewField69Field err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestField69Value(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField69Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField69Field err to be nil, got %v", err)
	}

	actual := field.Value()
	expected := []byte{0, 0, 0, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestField69String(t *testing.T) {
	data, _ := hex.DecodeString("000000450000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField69Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField69Field err to be nil, got %v", err)
	}

	actual := field.String()
	expected := "[0 0 0 0]"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
