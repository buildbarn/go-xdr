package opaque_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/buildbarn/go-xdr/pkg/compiler/tests/opaque"
	"github.com/stretchr/testify/require"
)

func TestFixedLengthOpaqueEncodedSize(t *testing.T) {
	for _, tt := range []struct {
		name             string
		encodedSizeBytes int
		writeTo          func(io.Writer) (int64, error)
		expected         []byte
	}{
		{
			name:             "BeforeBoundary",
			encodedSizeBytes: opaque.Opaque3EncodedSizeBytes,
			writeTo: func(w io.Writer) (int64, error) {
				return opaque.WriteOpaque3(w, &[3]byte{1, 2, 3})
			},
			expected: []byte{1, 2, 3, 0},
		},
		{
			name:             "OnBoundary",
			encodedSizeBytes: opaque.Opaque4EncodedSizeBytes,
			writeTo: func(w io.Writer) (int64, error) {
				return opaque.WriteOpaque4(w, &[4]byte{1, 2, 3, 4})
			},
			expected: []byte{1, 2, 3, 4},
		},
		{
			name:             "AfterBoundary",
			encodedSizeBytes: opaque.Opaque5EncodedSizeBytes,
			writeTo: func(w io.Writer) (int64, error) {
				return opaque.WriteOpaque5(w, &[5]byte{1, 2, 3, 4, 5})
			},
			expected: []byte{1, 2, 3, 4, 5, 0, 0, 0},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			n, err := tt.writeTo(buf)
			require.NoError(t, err)
			require.Equal(t, tt.expected, buf.Bytes())
			require.Equal(t, int64(len(tt.expected)), n)
			require.Equal(t, len(tt.expected), tt.encodedSizeBytes)
		})
	}
}
