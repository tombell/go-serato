package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewAlbumField(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	album, err := field.NewAlbumField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAlbumField err to be nil, got %v", err)
	}

	if album == nil {
		t.Fatal("expected album to not be nil")
	}
}

func TestNewAlbumFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F00720079000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewAlbumField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewAlbumField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewAlbumFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("000000090000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewAlbumField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewAlbumField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestAlbumValue(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	album, err := field.NewAlbumField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAlbumField err to be nil, got %v", err)
	}

	actual := album.Value()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestAlbumString(t *testing.T) {
	data, _ := hex.DecodeString("000000080000000C0047006C006F007200790000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	album, err := field.NewAlbumField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewAlbumField err to be nil, got %v", err)
	}

	actual := album.String()
	expected := "Glory"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
