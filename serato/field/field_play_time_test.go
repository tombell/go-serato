package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewPlayTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	playtime, err := field.NewPlayTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayTimeField err to be nil, got %v", err)
	}

	if playtime == nil {
		t.Fatal("expected playtime to not be nil")
	}
}

func TestNewPlayTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000002D0000000400009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewPlayTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewPlayTimeField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewPlayTimeFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000002E000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewPlayTimeField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewPlayTimeField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestPlayTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	playtime, err := field.NewPlayTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayTimeField err to be nil, got %v", err)
	}

	actual := playtime.Value()
	expected := 155

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestPlayTimeString(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000040000009B")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	playtime, err := field.NewPlayTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewPlayTimeField err to be nil, got %v", err)
	}

	actual := playtime.String()
	expected := "155"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
