package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewGroupingField(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	grouping, err := field.NewGroupingField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGroupingField err to be nil, got %v", err)
	}

	if grouping == nil {
		t.Fatal("expected grouping to not be nil")
	}
}

func TestNewGroupingFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E0074006900630073002000520072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewGroupingField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewGroupingField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewGroupingFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000140000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewGroupingField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewGroupingField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestGroupingValue(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	grouping, err := field.NewGroupingField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGroupingField err to be nil, got %v", err)
	}

	actual := grouping.Value()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestGroupingString(t *testing.T) {
	data, _ := hex.DecodeString("000000130000002400410074006C0061006E00740069006300730020005200650063006F0072006400730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	grouping, err := field.NewGroupingField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewGroupingField err to be nil, got %v", err)
	}

	actual := grouping.String()
	expected := "Atlantics Records"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
