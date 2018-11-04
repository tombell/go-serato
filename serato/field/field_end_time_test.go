package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewEndTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	endtime, err := field.NewEndTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewEndTimeField err to be nil, got %v", err)
	}

	if endtime == nil {
		t.Fatal("expected endtime to not be nil")
	}
}

func TestNewEndTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045BDA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewEndTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewEndTimeField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewEndTimeFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000002D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewEndTimeField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewEndTimeField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestEndTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	endtime, err := field.NewEndTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewEndTimeField err to be nil, got %v", err)
	}

	actual := endtime.Value()
	expected := time.Date(2018, time.September, 5, 20, 33, 39, 0, time.UTC)

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestEndTimeString(t *testing.T) {
	data, _ := hex.DecodeString("0000001D000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	endtime, err := field.NewEndTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewEndTimeField err to be nil, got %v", err)
	}

	actual := endtime.String()
	expected := "2018-09-05 20:33:39 +0000 UTC"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
