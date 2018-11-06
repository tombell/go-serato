package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

// XXX: Location field appears to always be empty in session files.

func TestNewLocationField(t *testing.T) {
	data, _ := hex.DecodeString("0000000300000072002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	location, err := field.NewLocationField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLocationField err to be nil, got %v", err)
	}

	if location == nil {
		t.Fatal("expected location to not be nil")
	}
}

func TestNewLocationFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("0000000300000072002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLocationField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewLocationField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewLocationFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000400000072002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewLocationField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewLocationField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestLocationValue(t *testing.T) {
	data, _ := hex.DecodeString("0000000300000072002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	location, err := field.NewLocationField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLocationField err to be nil, got %v", err)
	}

	actual := location.Value()
	expected := "/Users/tombell/Music/__ New __/Classic House Summer '18/"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestLocationString(t *testing.T) {
	data, _ := hex.DecodeString("0000000300000072002F00550073006500720073002F0074006F006D00620065006C006C002F004D0075007300690063002F005F005F0020004E006500770020005F005F002F0043006C0061007300730069006300200048006F007500730065002000530075006D006D006500720020002700310038002F0000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	location, err := field.NewLocationField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewLocationField err to be nil, got %v", err)
	}

	actual := location.String()
	expected := "/Users/tombell/Music/__ New __/Classic House Summer '18/"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
