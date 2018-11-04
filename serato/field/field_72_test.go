package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"reflect"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewField72Field(t *testing.T) {
	data, _ := hex.DecodeString("000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField72Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField72Field err to be nil, got %v", err)
	}

	if field == nil {
		t.Fatal("expected field to not be nil")
	}
}

func TestNewField72FieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000480000000400000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField72Field(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewField72Field err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewField72FieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000490000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewField72Field(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatal("expected NewField72Field err to be field.ErrUnexpectedIdentifier")
	}
}

func TestField72Value(t *testing.T) {
	data, _ := hex.DecodeString("000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField72Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField72Field err to be nil, got %v", err)
	}

	actual := field.Value()
	expected := []byte{0, 0, 0, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestField72String(t *testing.T) {
	data, _ := hex.DecodeString("000000480000000400000000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	field, err := field.NewField72Field(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewField72Field err to be nil, got %v", err)
	}

	actual := field.String()
	expected := "[0 0 0 0]"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
