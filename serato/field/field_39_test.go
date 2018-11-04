package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewField39Field(t *testing.T) {
	data, _ := hex.DecodeString("000000270000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField39Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField39Field err to be nil, got %v", err)
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField39FieldEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000002700000001")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField39Field(hdr, buf)
	if err != io.EOF {
		t.Fatal("expected NewField39Field err to be EOF")
	}
}

func TestNewField39FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000490000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField39Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewField39Field err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestField39Value(t *testing.T) {
	data, _ := hex.DecodeString("000000270000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField39Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField39Field err to be nil, got %v", err)
	}

	actual := field.Value()
	expected := []byte{1}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestField39String(t *testing.T) {
	data, _ := hex.DecodeString("000000270000000101")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField39Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField39Field err to be nil, got %v", err)
	}

	actual := field.String()
	expected := "[1]"

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
