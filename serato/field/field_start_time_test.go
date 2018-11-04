package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewStartTimeField(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	starttime, err := field.NewStartTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewStartTimeField err to be nil, got %v", err)
	}

	if starttime == nil {
		t.Fatal("expected starttime to not be nil")
	}
}

func TestNewStartTimeFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B908")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewStartTimeField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewStartTimeField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewStartTimeFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000002C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewStartTimeField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewStartTimeField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestStartTimeValue(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	starttime, err := field.NewStartTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewStartTimeField err to be nil, got %v", err)
	}

	actual := starttime.Value()
	expected := time.Date(2018, time.September, 5, 20, 31, 04, 0, time.UTC)

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestStartTimeString(t *testing.T) {
	data, _ := hex.DecodeString("0000001C000000045B903D08")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	starttime, err := field.NewStartTimeField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewStartTimeField err to be nil, got %v", err)
	}

	actual := starttime.String()
	expected := "2018-09-05 20:31:04 +0000 UTC"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
