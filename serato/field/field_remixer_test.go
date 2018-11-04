package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewRemixerField(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	remixer, err := field.NewRemixerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewRemixerField err to be nil, got %v", err)
	}

	if remixer == nil {
		t.Fatal("expected remixer to not be nil")
	}
}

func TestNewRemixerFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C00650073")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewRemixerField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewRemixerField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewRemixerFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000001500000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewRemixerField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewRemixerField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestRemixerValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	remixer, err := field.NewRemixerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewRemixerField err to be nil, got %v", err)
	}

	actual := remixer.Value()
	expected := "Frankie Knuckles"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestRemixerString(t *testing.T) {
	data, _ := hex.DecodeString("0000001400000022004600720061006E006B006900650020004B006E00750063006B006C006500730000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	remixer, err := field.NewRemixerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewRemixerField err to be nil, got %v", err)
	}

	actual := remixer.String()
	expected := "Frankie Knuckles"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
