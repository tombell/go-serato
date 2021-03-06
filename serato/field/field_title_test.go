package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewTitleField(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	title, err := field.NewTitleField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewTitleField err to be nil, got %v", err)
	}

	if title == nil {
		t.Fatal("expected title to not be nil")
	}
}

func TestNewTitleFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F007500200")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewTitleField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewTitleField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewTitleFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewTitleField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewTitleField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestTitleValue(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	title, err := field.NewTitleField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewTitleField err to be nil, got %v", err)
	}

	actual := title.Value()
	expected := "Do You Wanna House (Original Mix)"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestTitleString(t *testing.T) {
	data, _ := hex.DecodeString("00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	title, err := field.NewTitleField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewTitleField err to be nil, got %v", err)
	}

	actual := title.String()
	expected := "Do You Wanna House (Original Mix)"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
