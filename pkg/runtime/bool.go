package runtime

import (
	"errors"
	"io"
)

var (
	marshaledFalse = [4]byte{0, 0, 0, 0}
	marshaledTrue  = [4]byte{0, 0, 0, 1}
)

var errInvalidBoolean = errors.New("boolean value not zero or one")

// ReadBool reads an XDR encoded boolean value.
func ReadBool(r io.Reader) (bool, int64, error) {
	var b [4]byte
	n, err := io.ReadFull(r, b[:])
	if err != nil {
		return false, int64(n), err
	}
	switch b {
	case marshaledFalse:
		return false, int64(n), nil
	case marshaledTrue:
		return true, int64(n), nil
	default:
		return false, int64(n), errInvalidBoolean
	}
}

// WriteBool writes an XDR encoded boolean value.
func WriteBool(w io.Writer, v bool) (int64, error) {
	if v {
		n, err := w.Write(marshaledTrue[:])
		return int64(n), err
	}
	n, err := w.Write(marshaledFalse[:])
	return int64(n), err
}
