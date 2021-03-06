package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewUpdatedAtField(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	updatedat, err := field.NewUpdatedAtField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUpdatedAtField err to be nil, got %v", err)
	}

	if updatedat == nil {
		t.Fatal("expected updatedat to not be nil")
	}
}

func TestNewUpdatedAtFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewUpdatedAtField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewUpdatedAtField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewUpdatedAtFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000045000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewUpdatedAtField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewUpdatedAtField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestUpdatedAtValue(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	updatedat, err := field.NewUpdatedAtField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUpdatedAtField err to be nil, got %v", err)
	}

	actual := updatedat.Value()
	expected := time.Date(2018, time.September, 5, 20, 33, 39, 0, time.UTC)

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestUpdatedAtString(t *testing.T) {
	data, _ := hex.DecodeString("00000035000000045B903DA3")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	updatedat, err := field.NewUpdatedAtField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUpdatedAtField err to be nil, got %v", err)
	}

	actual := updatedat.String()
	expected := "2018-09-05 20:33:39 +0000 UTC"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
