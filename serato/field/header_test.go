package field_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/field"
)

func TestNewHeader(t *testing.T) {
	data := generateBytes(t, "0000000100000004")
	buf := bytes.NewBuffer(data)

	hdr, err := field.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	expectedIdentifier := uint32(1)
	if hdr.Identifier != expectedIdentifier {
		t.Fatalf("expected identifier to be %v", expectedIdentifier)
	}

	expectedLength := uint32(4)
	if hdr.Length != expectedLength {
		t.Fatalf("expected length to be %v", expectedLength)
	}
}

func TestNewHeaderUnexpectedEOF(t *testing.T) {
	data := generateBytes(t, "000102")
	buf := bytes.NewBuffer(data)

	_, err := field.NewHeader(buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected err to be unexpected eof error")
	}
}

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"artist", generateBytes(t, "00000007000000360044004A0020004600610076006F0072006900740065002C00200044004A0020004B00680061007200690074006F006E006F00760000"), "field (7) length (54)"},
		{"title", generateBytes(t, "00000006000000440044006F00200059006F0075002000570061006E006E006100200048006F00750073006500200028004F0072006900670069006E0061006C0020004D0069007800290000"), "field (6) length (68)"},
		{"starttime", generateBytes(t, "0000001C000000045B903D08"), "field (28) length (4)"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := field.NewHeader(buf)
			if err != nil {
				t.Fatal("expected err to be nil")
			}

			actual := hdr.String()
			if actual != tc.expected {
				t.Fatalf("expected string to be %v, got %v", tc.expected, actual)
			}
		})
	}
}
