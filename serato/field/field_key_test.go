package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewKeyField(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	key, err := field.NewKeyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewKeyField err to be nil, got %v", err)
	}

	if key == nil {
		t.Fatal("expected key to not be nil")
	}
}

func TestNewKeyFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewKeyField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewKeyField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewKeyFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000043000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewKeyField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewKeyField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestKeyValue(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	key, err := field.NewKeyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewKeyField err to be nil, got %v", err)
	}

	actual := key.Value()
	expected := "Cm"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestKeyString(t *testing.T) {
	data, _ := hex.DecodeString("00000033000000060043006D0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	key, err := field.NewKeyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewKeyField err to be nil, got %v", err)
	}

	actual := key.String()
	expected := "Cm"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
