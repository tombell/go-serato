package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewSessionIDField(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	sessionid, err := field.NewSessionIDField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewSessionIDField err to be nil, got %v", err)
	}

	if sessionid == nil {
		t.Fatal("expected sessionid to not be nil")
	}
}

func TestNewSessionIDFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000002")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewSessionIDField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewSessionIDField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewSessionIDFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000004000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewSessionIDField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewSessionIDField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestSessionIDValue(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	sessionid, err := field.NewSessionIDField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewSessionIDField err to be nil, got %v", err)
	}

	actual := sessionid.Value()
	expected := 210

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestSessionIDString(t *testing.T) {
	data, _ := hex.DecodeString("0000003000000004000000D2")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	sessionid, err := field.NewSessionIDField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewSessionIDField err to be nil, got %v", err)
	}

	actual := sessionid.String()
	expected := "210"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
