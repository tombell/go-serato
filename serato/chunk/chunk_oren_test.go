package chunk_test

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/chunk"
)

func TestNewOrenChunk(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewOrenChunk err to be nil, got %v", err)
	}

	if oren == nil {
		t.Fatal("expected oren to not be nil")
	}
}

func TestNewOrenChunkUnexpectedEOF(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = chunk.NewOrenChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("expected NewOrenChunk err to be io.ErrUnexpectedEOF, got %v", err)
	}
}

func TestNewOrenChunkUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("6E72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = chunk.NewOrenChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewOrenChunk err to be chunk.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestNewOrenChunkUentUnexpectedIdentifier(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C74656E73000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	_, err = chunk.NewOrenChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatalf("expected NewOrenChunk err to be chunk.ErrUnexpectedIdentifier, got %v", err)
	}
}

func TestOrenHeader(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewOrenChunk err to be nil, got %v", err)
	}

	if oren.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestOrenType(t *testing.T) {
	data, _ := hex.DecodeString("6F72656E0000000C75656E74000000040000000F")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatalf("expected NewHeader err to be nil, got %v", err)
	}

	oren, err := chunk.NewOrenChunk(hdr, buf)
	if err != nil {
		t.Fatalf("expected NewOrenChunk err to be nil, got %v", err)
	}

	actual := oren.Type()
	expected := "oren"

	if actual != expected {
		t.Fatalf("expected type to be %v, got %v", expected, actual)
	}
}
