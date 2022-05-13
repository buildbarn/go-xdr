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

func TestReadBool(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		r.EXPECT().Read(gomock.Any()).Return(0, errors.New("disk failure"))

		_, n, err := runtime.ReadBool(r)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("EOF", func(t *testing.T) {
		_, n, err := runtime.ReadBool(bytes.NewBuffer(nil))
		require.Equal(t, io.EOF, err)
		require.Equal(t, int64(0), n)
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		input := [...]byte{1, 2, 3}
		for i := 1; i <= len(input); i++ {
			_, n, err := runtime.ReadBool(bytes.NewBuffer(input[:i]))
			require.Equal(t, io.ErrUnexpectedEOF, err)
			require.Equal(t, int64(i), n)
		}
	})

	t.Run("InvalidValue", func(t *testing.T) {
		_, n, err := runtime.ReadBool(bytes.NewBuffer([]byte{0, 0, 0, 2}))
		require.Equal(t, errors.New("boolean value not zero or one"), err)
		require.Equal(t, int64(4), n)
	})

	t.Run("False", func(t *testing.T) {
		v, n, err := runtime.ReadBool(bytes.NewBuffer([]byte{0, 0, 0, 0}))
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.False(t, v)
	})

	t.Run("True", func(t *testing.T) {
		v, n, err := runtime.ReadBool(bytes.NewBuffer([]byte{0, 0, 0, 1}))
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.True(t, v)
	})
}

func TestWriteBool(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, errors.New("disk failure"))

		n, err := runtime.WriteBool(w, true)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("False", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteBool(b, false)
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, []byte{0, 0, 0, 0}, b.Bytes())
	})

	t.Run("True", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteBool(b, true)
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, []byte{0, 0, 0, 1}, b.Bytes())
	})
}
