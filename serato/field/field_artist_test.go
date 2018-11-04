package field_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewArtistField(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewArtistField err to be nil, got %v", err)
	}

	if artist == nil {
		t.Fatal("expected artist to not be nil")
	}
}

func TestNewArtistFieldUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A002000460061007600")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewArtistField(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewArtistField err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewArtistFieldUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("00000005000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = field.NewArtistField(hdr, buf)
	if err != field.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewArtistField err to be field.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestArtistValue(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewArtistField err to be nil, got %v", err)
	}

	actual := artist.Value()
	expected := "DJ Favorite, DJ Kharitonov"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}

func TestArtistString(t *testing.T) {
	data, _ := hex.DecodeString("00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	artist, err := field.NewArtistField(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewArtistField err to be nil, got %v", err)
	}

	actual := artist.String()
	expected := "DJ Favorite, DJ Kharitonov"

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
