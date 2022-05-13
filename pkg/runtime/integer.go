package runtime

import (
	"encoding/binary"
	"io"
)

// ReadInt reads an XDR encoded 32-bit signed integer value.
func ReadInt(r io.Reader) (int32, int64, error) {
	var b [4]byte
	n, err := io.ReadFull(r, b[:])
	return int32(binary.BigEndian.Uint32(b[:])), int64(n), err
}

// WriteInt writes an XDR encoded 32-bit signed integer value.
func WriteInt(w io.Writer, i int32) (int64, error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(i))
	n, err := w.Write(b[:])
	return int64(n), err
}

// ReadUnsignedInt reads an XDR encoded 32-bit unsigned integer value.
func ReadUnsignedInt(r io.Reader) (uint32, int64, error) {
	var b [4]byte
	n, err := io.ReadFull(r, b[:])
	return binary.BigEndian.Uint32(b[:]), int64(n), err
}

// WriteUnsignedInt writes an XDR encoded 32-bit unsigned integer value.
func WriteUnsignedInt(w io.Writer, i uint32) (int64, error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], i)
	n, err := w.Write(b[:])
	return int64(n), err
}

// ReadHyper reads an XDR encoded 64-bit signed integer value.
func ReadHyper(r io.Reader) (int64, int64, error) {
	var b [8]byte
	n, err := io.ReadFull(r, b[:])
	return int64(binary.BigEndian.Uint64(b[:])), int64(n), err
}

// WriteHyper writes an XDR encoded 64-bit signed integer value.
func WriteHyper(w io.Writer, i int64) (int64, error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(i))
	n, err := w.Write(b[:])
	return int64(n), err
}

// ReadUnsignedHyper reads an XDR encoded 64-bit unsigned integer value.
func ReadUnsignedHyper(r io.Reader) (uint64, int64, error) {
	var b [8]byte
	n, err := io.ReadFull(r, b[:])
	return binary.BigEndian.Uint64(b[:]), int64(n), err
}

// WriteUnsignedHyper writes an XDR encoded 64-bit unsigned integer value.
func WriteUnsignedHyper(w io.Writer, i uint64) (int64, error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], i)
	n, err := w.Write(b[:])
	return int64(n), err
}
