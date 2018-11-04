package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewComposerField(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	composer, err := field.NewComposerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewComposerField err to be nil, got %v", err)
	}

	if composer == nil {
		t.Fatal("expected composer to not be nil")
	}
}

func TestNewComposerFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F0072007900")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewComposerField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewComposerField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewComposerFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000150000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewComposerField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewComposerField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestComposerValue(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	composer, err := field.NewComposerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewComposerField err to be nil, got %v", err)
	}

	actual := composer.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestComposerString(t *testing.T) {
	data, _ := hex.DecodeString("000000160000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	composer, err := field.NewComposerField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewComposerField err to be nil, got %v", err)
	}

	actual := composer.String()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
