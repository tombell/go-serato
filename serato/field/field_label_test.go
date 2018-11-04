package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewLabelField(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	label, err := field.NewLabelField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLabelField err to be nil, got %v", err)
	}

	if label == nil {
		t.Fatal("expected label to not be nil")
	}
}

func TestNewLabelFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLabelField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewLabelField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewLabelFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000140000002400410074006C0061006E00740069006300730020005200650063006F0000730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLabelField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewLabelField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestLabelValue(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	label, err := field.NewLabelField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLabelField err to be nil, got %v", err)
	}

	actual := label.Value()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestLabelString(t *testing.T) {
	data, _ := hex.DecodeString("000000150000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	label, err := field.NewLabelField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLabelField err to be nil, got %v", err)
	}

	actual := label.String()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
