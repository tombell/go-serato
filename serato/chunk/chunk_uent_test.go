package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/chunk"
)

func TestNewUentChunk(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUentChunk err to be nil, got %v", err)
	}

	if uent == nil {
		t.Fatal("expected uent to not be nil")
	}
}

func TestNewUentChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("75656E7400000004000000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = chunk.NewUentChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewUentChunk err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewUentChunkUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("74656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = chunk.NewUentChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewUentChunk err to be chunk.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestUentHeader(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUentChunk err to be nil, got %v", err)
	}

	if uent.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestUentType(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUentChunk err to be nil, got %v", err)
	}

	actual := uent.Type()
	expected := "uent"

	if actual != expected {
		t.Fatalf("expected type to be %v, got %v", expected, actual)
	}
}

func TestUentValue(t *testing.T) {
	data, _ := hex.DecodeString("75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	uent, err := chunk.NewUentChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewUentChunk err to be nil, got %v", err)
	}

	actual := uent.Value()
	expected := 15

	if actual != expected {
		t.Fatalf("expected value to be %v, got %v", expected, actual)
	}
}
