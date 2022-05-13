package runtime

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

func validateASCII(s string) error {
	for i := 0; i < len(s); i++ {
		if c := s[i]; c > unicode.MaxASCII {
			return fmt.Errorf("character at offset %d has non-ASCII value %#x", i, c)
		}
	}
	return nil
}

// ReadASCIIString reads an XDR encoded ASCII string value.
func ReadASCIIString(r io.Reader, maximumSizeBytes uint32) (string, int64, error) {
	var builder strings.Builder
	n, err := copyVariableLengthOpaque(r, &builder, maximumSizeBytes)
	if err != nil {
		return "", n, err
	}
	s := builder.String()
	if err := validateASCII(s); err != nil {
		return "", n, err
	}
	return s, n, err
}

// WriteASCIIString writes an XDR encoded ASCII string value.
func WriteASCIIString(w io.Writer, maximumSizeBytes uint32, s string) (int64, error) {
	if err := validateASCII(s); err != nil {
		return 0, err
	}
	return WriteVariableLengthOpaque(w, maximumSizeBytes, []byte(s))
}
