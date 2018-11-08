package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

// XXX: Frequency field appears to always be empty in session files.

func TestNewFrequencyField(t *testing.T) {
	data, _ := hex.DecodeString("0000000E0000000C00340034002E0031006B0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	frequency, err := field.NewFrequencyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFrequencyField err to be nil, got %v", err)
	}

	if frequency == nil {
		t.Fatal("expected frequency to not be nil")
	}
}

func TestNewFrequencyFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000E0000000C00340034002E0031006B00")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewFrequencyField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewFrequencyField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewFrequencyFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000F0000000C00340034002E0031006B0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewFrequencyField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewFrequencyField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestFrequencyValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000E0000000C00340034002E0031006B0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	frequency, err := field.NewFrequencyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFrequencyField err to be nil, got %v", err)
	}

	actual := frequency.Value()
	expected := "44.1k"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestFrequencyString(t *testing.T) {
	data, _ := hex.DecodeString("0000000E0000000C00340034002E0031006B0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	frequency, err := field.NewFrequencyField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFrequencyField err to be nil, got %v", err)
	}

	actual := frequency.String()
	expected := "44.1k"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
