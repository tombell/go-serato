package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

// XXX: Filename field appears to always be empty in session files.

func TestNewFilenameField(t *testing.T) {
	data, _ := hex.DecodeString("0000000400000059310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	filename, err := field.NewFilenameField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFilenameField err to be nil, got %v", err)
	}

	if filename == nil {
		t.Fatal("expected filename to not be nil")
	}
}

func TestNewFilenameFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000040000005A00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D00700033")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewFilenameField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewFilenameField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewFilenameFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("0000000500000059310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewFilenameField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewFilenameField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestFilenameValue(t *testing.T) {
	data, _ := hex.DecodeString("000000040000005A00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	filename, err := field.NewFilenameField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFilenameField err to be nil, got %v", err)
	}

	actual := filename.Value()
	expected := "10947360_Do_You_Wanna_House_Original_Mix.mp3"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestFilenameString(t *testing.T) {
	data, _ := hex.DecodeString("000000040000005A00310030003900340037003300360030005F0044006F005F0059006F0075005F00570061006E006E0061005F0048006F007500730065005F004F0072006900670069006E0061006C005F004D00690078002E006D007000330000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	filename, err := field.NewFilenameField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewFilenameField err to be nil, got %v", err)
	}

	actual := filename.String()
	expected := "10947360_Do_You_Wanna_House_Original_Mix.mp3"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
