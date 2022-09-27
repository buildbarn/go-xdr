package runtime

import (
	"bytes"
	"fmt"
	"io"
)

func readPadding(r io.Reader, nTotal int64) (int64, error) {
	paddingToSkip := nTotal % 4
	if paddingToSkip == 0 {
		return nTotal, nil
	}

	var padding [4]byte
	paddingSlice := padding[paddingToSkip:]
	nPadding, err := io.ReadFull(r, paddingSlice)
	nTotal += int64(nPadding)
	if err != nil {
		if err == io.EOF {
			return nTotal, io.ErrUnexpectedEOF
		}
		return nTotal, err
	}
	if padding != [4]byte{} {
		return nTotal, fmt.Errorf("padding contains non-zero values %v", paddingSlice)
	}
	return nTotal, nil
}

var zeroPadding [4]byte

func writePadding(w io.Writer, n int64) (int64, error) {
	paddingToSkip := n % 4
	if paddingToSkip == 0 {
		return n, nil
	}
	nPadding, err := w.Write(zeroPadding[paddingToSkip:])
	return n + int64(nPadding), err
}

// ReadFixedLengthOpaque reads an XDR encoded fixed length opaque value.
func ReadFixedLengthOpaque(r io.Reader, b []byte) (int64, error) {
	nData, err := io.ReadFull(r, b)
	if err != nil {
		return int64(nData), err
	}
	return readPadding(r, int64(nData))
}

// WriteFixedLengthOpaque writes an XDR encoded fixed length opaque value.
func WriteFixedLengthOpaque(w io.Writer, b []byte) (int64, error) {
	nData, err := w.Write(b)
	if err != nil {
		return int64(nData), err
	}
	return writePadding(w, int64(nData))
}

func copyVariableLengthOpaque(r io.Reader, w io.Writer, maximumSizeBytes uint32) (int64, error) {
	length, nTotal, err := ReadUnsignedInt(r)
	if err != nil || length == 0 {
		return nTotal, err
	}
	if length > maximumSizeBytes {
		return nTotal, fmt.Errorf("size of %d bytes exceeds field's maximum of %d bytes", length, maximumSizeBytes)
	}

	nCopied, err := io.CopyN(w, r, int64(length))
	nTotal += nCopied
	if err != nil {
		if err == io.EOF {
			return nTotal, io.ErrUnexpectedEOF
		}
		return nTotal, err
	}
	return readPadding(r, nTotal)
}

// ReadVariableLengthOpaque reads an XDR encoded variable length opaque value.
func ReadVariableLengthOpaque(r io.Reader, maximumSizeBytes uint32) ([]byte, int64, error) {
	b := bytes.NewBuffer(nil)
	n, err := copyVariableLengthOpaque(r, b, maximumSizeBytes)
	if err != nil {
		return nil, n, err
	}
	return b.Bytes(), n, nil
}

// WriteVariableLengthOpaque writes an XDR encoded variable length opaque value.
func WriteVariableLengthOpaque(w io.Writer, maximumSizeBytes uint32, b []byte) (int64, error) {
	if uint64(len(b)) > uint64(maximumSizeBytes) {
		return 0, fmt.Errorf("size of %d bytes exceeds field's maximum of %d bytes", len(b), maximumSizeBytes)
	}

	nTotal, err := WriteUnsignedInt(w, uint32(len(b)))
	if err != nil {
		return nTotal, err
	}

	nString, err := w.Write(b)
	nTotal += int64(nString)
	if err != nil {
		return nTotal, err
	}
	return writePadding(w, nTotal)
}
