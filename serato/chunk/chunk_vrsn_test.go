package chunk_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/tombell/go-serato/serato/chunk"
)

func TestNewVrsnChunk(t *testing.T) {
	data := generateBytes(t, "7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewVrsnChunk err to be nil")
	}

	if vrsn == nil {
		t.Fatal("expected vrsn to not be nil")
	}
}

func TestNewVrsnChunkUnexpectedEOF(t *testing.T) {
	data := generateBytes(t, "7672736E00000037600450020005200650076006")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewVrsnChunk(hdr, buf)
	if err != io.ErrUnexpectedEOF {
		t.Fatal("expected NewVrsnChunk err to be ErrUnexpectedEOF")
	}
}

func TestNewVrsnChunkUnexpectedIdentifier(t *testing.T) {
	data := generateBytes(t, "7572736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	_, err = chunk.NewVrsnChunk(hdr, buf)
	if err != chunk.ErrUnexpectedIdentifier {
		t.Fatal("expected NewVrsnChunk err to be ErrUnexpectedIdentifier")
	}
}

func TestVrsnHeader(t *testing.T) {
	data := generateBytes(t, "7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewVrsnChunk err to be nil")
	}

	if vrsn.Header() != hdr {
		t.Fatal("expected header to be the same")
	}
}

func TestVrsnType(t *testing.T) {
	data := generateBytes(t, "7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewVrsnChunk err to be nil")
	}

	actual := vrsn.Type()
	expected := "vrsn"

	if actual != expected {
		t.Fatalf("expected type to be %v, got %v", expected, actual)
	}
}

func TestVrsnVersion(t *testing.T) {
	data := generateBytes(t, "7672736E0000003C0031002E0030002F00530065007200610074006F002000530063007200610074006300680020004C0049005600450020005200650076006900650077")
	buf := bytes.NewBuffer(data)

	hdr, err := chunk.NewHeader(buf)
	if err != nil {
		t.Fatal("expected NewHeader err to be nil")
	}

	vrsn, err := chunk.NewVrsnChunk(hdr, buf)
	if err != nil {
		t.Fatal("expected NewVrsnChunk err to be nil")
	}

	actual := vrsn.Version()
	expected := "1.0/Serato Scratch LIVE Review"

	if actual != expected {
		t.Fatalf("expected version to be %v, got %v", expected, actual)
	}
}
