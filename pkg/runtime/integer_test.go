package runtime_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/buildbarn/go-xdr/internal/mock"
	"github.com/buildbarn/go-xdr/pkg/runtime"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestReadInt(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		r.EXPECT().Read(gomock.Any()).Return(0, errors.New("disk failure"))

		_, n, err := runtime.ReadInt(r)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("EOF", func(t *testing.T) {
		_, n, err := runtime.ReadInt(bytes.NewBuffer(nil))
		require.Equal(t, io.EOF, err)
		require.Equal(t, int64(0), n)
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		input := [...]byte{1, 2, 3}
		for i := 1; i <= len(input); i++ {
			_, n, err := runtime.ReadInt(bytes.NewBuffer(input[:i]))
			require.Equal(t, io.ErrUnexpectedEOF, err)
			require.Equal(t, int64(i), n)
		}
	})

	t.Run("Success", func(t *testing.T) {
		v, n, err := runtime.ReadInt(bytes.NewBuffer([]byte{0x2d, 0x11, 0x8b, 0xa4}))
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, int32(0x2d118ba4), v)
	})
}

func TestWriteInt(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, errors.New("disk failure"))

		n, err := runtime.WriteInt(w, 0x6e3895ee)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("Success", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteInt(b, 0x7475cd45)
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, []byte{0x74, 0x75, 0xcd, 0x45}, b.Bytes())
	})
}
