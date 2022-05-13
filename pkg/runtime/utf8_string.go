package runtime

import (
	"errors"
	"io"
	"strings"
	"unicode/utf8"
)

var errBadUTF8 = errors.New("string is not valid UTF-8")

// ReadUTF8String reads an XDR encoded UTF-8 string value.
//
// UTF-8 strings aren't part of RFC 4506, but are used by protocols such
// as NFSv4, that it warrants supporting it as a native type.
func ReadUTF8String(r io.Reader, maximumSizeBytes uint32) (string, int64, error) {
	var builder strings.Builder
	n, err := copyVariableLengthOpaque(r, &builder, maximumSizeBytes)
	if err != nil {
		return "", n, err
	}
	s := builder.String()
	if !utf8.ValidString(s) {
		return "", n, errBadUTF8
	}
	return s, n, err
}

// WriteUTF8String writes an XDR encoded UTF-8 string value.
//
// UTF-8 strings aren't part of RFC 4506, but are used by protocols such
// as NFSv4, that it warrants supporting it as a native type.
func WriteUTF8String(w io.Writer, maximumSizeBytes uint32, s string) (int64, error) {
	if !utf8.ValidString(s) {
		return 0, errBadUTF8
	}
	return WriteVariableLengthOpaque(w, maximumSizeBytes, []byte(s))
}
