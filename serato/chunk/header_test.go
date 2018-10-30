package chunk_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/chunk"
)

func TestNewHeader(t *testing.T) {
	data := generateBytes(t, "7672736E0000003C")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	expectedIdentifier := [4]byte{0x76, 0x72, 0x73, 0x6E}
	if hdr.Identifier != expectedIdentifier {
		t.Fatalf("expected identifier to be %v", expectedIdentifier)
	}

	expectedLength := uint32(60)
	if hdr.Length != expectedLength {
		t.Fatalf("expected length to be %v", expectedLength)
	}
}

func TestNewHeaderUnexpectedEOF(t *testing.T) {
	data := generateBytes(t, "000102")
	buf := bytes.NewBuffer(data)

	_, err := chunk.NewHeader(buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected err to be io.ErrUnexpectedEOF")
	}
}

func TestHeaderType(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		expectedType string
	}{
		{"vrsn", generateBytes(t, "7672736E0000003C"), "vrsn"},
		{"oent", generateBytes(t, "6F656E740000028F"), "oent"},
		{"adat", generateBytes(t, "6164617400000287"), "adat"},
		{"oren", generateBytes(t, "6F72656E000002EF"), "oren"},
		{"uent", generateBytes(t, "75656E74000001AE"), "uent"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := chunk.NewHeader(buf)
			if err != nil {
				t.Fatal("expected err to be nil")
			}

			actual := hdr.Type()
			if actual != tc.expectedType {
				t.Fatalf("expected type to be %v, got %v", tc.expectedType, actual)
			}
		})
	}
}

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"vrsn", generateBytes(t, "7672736E0000003C"), "chunk (vrsn) length (60)"},
		{"oent", generateBytes(t, "6F656E740000028F"), "chunk (oent) length (655)"},
		{"adat", generateBytes(t, "6164617400000287"), "chunk (adat) length (647)"},
		{"oren", generateBytes(t, "6F72656E0000000C"), "chunk (oren) length (12)"},
		{"uent", generateBytes(t, "75656E7400000008"), "chunk (uent) length (8)"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.input)

			hdr, err := chunk.NewHeader(buf)
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
