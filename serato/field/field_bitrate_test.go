package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

// XXX: Bitrate field appears to always be empty in session files.

func TestNewBitrateField(t *testing.T) {
	data, _ := hex.DecodeString("0000000D00000014003300320030002E0030006B0062007000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bitrate, err := field.NewBitrateField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBitrateField err to be nil, got %v", err)
	}

	if bitrate == nil {
		t.Fatal("expected bitrate to not be nil")
	}
}

func TestNewBitrateFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000D00000014003300320030002E0030006B00620070007300")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewBitrateField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewBitrateField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewBitrateFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000E00000014003300320030002E0030006B0062007000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewBitrateField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewBitrateField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestBitrateValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000D00000014003300320030002E0030006B0062007000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bitrate, err := field.NewBitrateField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBitrateField err to be nil, got %v", err)
	}

	actual := bitrate.Value()
	expected := "320.0kbps"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestBitrateString(t *testing.T) {
	data, _ := hex.DecodeString("0000000D00000014003300320030002E0030006B0062007000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bitrate, err := field.NewBitrateField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBitrateField err to be nil, got %v", err)
	}

	actual := bitrate.String()
	expected := "320.0kbps"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
