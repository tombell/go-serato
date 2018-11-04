package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewCommentField(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	comment, err := field.NewCommentField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewCommentField err to be nil, got %v", err)
	}

	if comment == nil {
		t.Fatal("expected comment to not be nil")
	}
}

func TestNewCommentFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F0072007900")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewCommentField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewCommentField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewCommentFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000120000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewCommentField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewCommentField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestCommentValue(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	comment, err := field.NewCommentField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewCommentField err to be nil, got %v", err)
	}

	actual := comment.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestCommentString(t *testing.T) {
	data, _ := hex.DecodeString("000000110000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	comment, err := field.NewCommentField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewCommentField err to be nil, got %v", err)
	}

	actual := comment.String()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
