package rfc4506_test

import (
	"bytes"
	"testing"

	"github.com/buildbarn/go-xdr/pkg/compiler/tests/rfc4506"
	"github.com/stretchr/testify/require"
)

func TestRFC4506Chapter7(t *testing.T) {
	// Example taken from RFC 4506, chapter 7.
	m := rfc4506.File{
		Filename: "sillyprog",
		Type: &rfc4506.Filetype_EXEC{
			Interpretor: "lisp",
		},
		Owner: "john",
		Data:  []byte("(quit)"),
	}
	b := []byte{
		0x00, 0x00, 0x00, 0x09, // length of filename = 9
		0x73, 0x69, 0x6c, 0x6c, // filename characters
		0x79, 0x70, 0x72, 0x6f, // ... and more characters ...
		0x67, 0x00, 0x00, 0x00, // ... and 3 zero-bytes of fill
		0x00, 0x00, 0x00, 0x02, // filekind is EXEC = 2
		0x00, 0x00, 0x00, 0x04, // length of interpretor = 4
		0x6c, 0x69, 0x73, 0x70, // interpretor characters
		0x00, 0x00, 0x00, 0x04, // length of owner = 4
		0x6a, 0x6f, 0x68, 0x6e, // owner characters
		0x00, 0x00, 0x00, 0x06, // length of file data = 6
		0x28, 0x71, 0x75, 0x69, // file data bytes ...
		0x74, 0x29, 0x00, 0x00, // ... and 2 zero-bytes of fill
	}

	t.Run("ReadFrom", func(t *testing.T) {
		var m2 rfc4506.File
		n, err := m2.ReadFrom(bytes.NewBuffer(b))
		require.NoError(t, err)
		require.Equal(t, int64(len(b)), n)
		require.Equal(t, m, m2)
	})

	t.Run("WriteTo", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		n, err := m.WriteTo(buf)
		require.NoError(t, err)
		require.Equal(t, int64(len(b)), n)
		require.Equal(t, b, buf.Bytes())
	})
}
