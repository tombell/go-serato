package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewBPMField(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bpm, err := field.NewBPMField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBPMField err to be nil, got %v", err)
	}

	if bpm == nil {
		t.Fatal("expected bpm to not be nil")
	}
}

func TestNewBPMFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000F000000040000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewBPMField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewBPMField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewBPMFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000D0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewBPMField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewBPMField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestBPMValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bpm, err := field.NewBPMField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBPMField err to be nil, got %v", err)
	}

	actual := bpm.Value()
	expected := 119

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestBPMString(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000400000077")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	bpm, err := field.NewBPMField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewBPMField err to be nil, got %v", err)
	}

	actual := bpm.String()
	expected := "119"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
